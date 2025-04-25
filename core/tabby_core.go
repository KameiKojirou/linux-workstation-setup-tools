package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/charmbracelet/log"
	"slices"
)

func InstallTabby() {
	log.Info("Installing Tabby terminal...")

	// Check if user is in the required groups
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Failed to get current user", "error", err)
	}

	// Get current user groups
	cmd := exec.Command("groups", currentUser.Username)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("Failed to get user groups", "error", err)
	}

	userGroups := strings.Fields(string(output))
	
	// Check and add missing groups
	requiredGroups := []string{"tty", "dialout"}
	for _, group := range requiredGroups {
		if !slices.Contains(userGroups, group) {
			log.Info("Adding user to group", "group", group)
			addCmd := exec.Command("sudo", "usermod", "-a", "-G", group, currentUser.Username)
			if err := addCmd.Run(); err != nil {
				log.Error("Failed to add user to group", "group", group, "error", err)
			}
		}
	}

	// Get latest Tabby release
	latestDeb, err := getLatestTabbyDeb()
	if err != nil {
		log.Fatal("Failed to get latest Tabby release", "error", err)
	}

	// Download the .deb file
	debPath := filepath.Join(os.TempDir(), "tabby-latest.deb")
	log.Info("Downloading Tabby", "url", latestDeb)
	if err := downloadFile(debPath, latestDeb); err != nil {
		log.Fatal("Failed to download Tabby", "error", err)
	}

	// Install the .deb file
	log.Info("Installing Tabby package...")
	installCmd := exec.Command("sudo", "dpkg", "-i", debPath)
	if err := installCmd.Run(); err != nil {
		log.Error("Failed to install Tabby package, attempting to fix dependencies", "error", err)
		fixCmd := exec.Command("sudo", "apt-get", "install", "-f", "-y")
		if err := fixCmd.Run(); err != nil {
			log.Fatal("Failed to fix dependencies", "error", err)
		}
	}

	// Clean up
	os.Remove(debPath)
	log.Info("Tabby installed successfully!")
}

func UpgradeTabby() {
	log.Info("Updating Tabby terminal...")
	
	// Get latest Tabby release
	latestDeb, err := getLatestTabbyDeb()
	if err != nil {
		log.Fatal("Failed to get latest Tabby release", "error", err)
	}

	// Download the .deb file
	debPath := filepath.Join(os.TempDir(), "tabby-latest.deb")
	log.Info("Downloading Tabby update", "url", latestDeb)
	if err := downloadFile(debPath, latestDeb); err != nil {
		log.Fatal("Failed to download Tabby update", "error", err)
	}

	// Install the .deb file to update
	log.Info("Installing Tabby update...")
	installCmd := exec.Command("sudo", "dpkg", "-i", debPath)
	if err := installCmd.Run(); err != nil {
		log.Error("Failed to install Tabby update, attempting to fix dependencies", "error", err)
		fixCmd := exec.Command("sudo", "apt-get", "install", "-f", "-y")
		if err := fixCmd.Run(); err != nil {
			log.Fatal("Failed to fix dependencies", "error", err)
		}
	}

	// Clean up
	os.Remove(debPath)
	log.Info("Tabby updated successfully!")
}

func UninstallTabby() {
	log.Info("Uninstalling Tabby terminal...")
	
	// Remove the package
	cmd := exec.Command("sudo", "apt-get", "remove", "-y", "tabby")
	if err := cmd.Run(); err != nil {
		log.Error("Failed to uninstall Tabby", "error", err)
	}
	
	// Remove any leftover configurations
	purgeCmd := exec.Command("sudo", "apt-get", "purge", "-y", "tabby")
	if err := purgeCmd.Run(); err != nil {
		log.Error("Failed to purge Tabby configurations", "error", err)
	}
	
	// Clean up
	cleanCmd := exec.Command("sudo", "apt-get", "autoremove", "-y")
	if err := cleanCmd.Run(); err != nil {
		log.Error("Failed to clean up dependencies", "error", err)
	}
	
	log.Info("Tabby uninstalled successfully!")
}

// Helper function to get the latest Tabby .deb URL from GitHub releases
func getLatestTabbyDeb() (string, error) {
	// Make request to GitHub API for latest release
	resp, err := http.Get("https://api.github.com/repos/Eugeny/tabby/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch releases: status code %d", resp.StatusCode)
	}

	// Parse the response
	var release struct {
		TagName string `json:"tag_name"`
		Assets []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
		} `json:"assets"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}

	// Get the version number without the 'v' prefix
	version := strings.TrimPrefix(release.TagName, "v")
	
	// Determine the architecture suffix to look for
	archSuffix := "linux-x64"
	if runtime.GOARCH == "arm64" {
		archSuffix = "linux-arm64"
	}

	// Construct the expected filename pattern
	expectedFilename := fmt.Sprintf("tabby-%s-%s.deb", version, archSuffix)

	// Find the matching asset
	for _, asset := range release.Assets {
		if asset.Name == expectedFilename {
			return asset.BrowserDownloadURL, nil
		}
	}

	// Fallback: try a more flexible search if exact match wasn't found
	for _, asset := range release.Assets {
		if strings.HasPrefix(asset.Name, "tabby-") && 
		   strings.Contains(asset.Name, archSuffix) && 
		   strings.HasSuffix(asset.Name, ".deb") {
			return asset.BrowserDownloadURL, nil
		}
	}

	return "", fmt.Errorf("couldn't find Tabby .deb package for %s architecture", runtime.GOARCH)
}

// Helper function to download a file from a URL
func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
