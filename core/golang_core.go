package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func InstallGolang() {
	log.Println("Installing Golang")
	// Step 0: Check if Golang is already installed.
	checkCmd := exec.Command("sh", "-c", "go version")
	checkCmd.Stdout = os.Stdout
	checkCmd.Stderr = os.Stderr
	if err := checkCmd.Run(); err == nil {
		log.Println("Golang is already installed. Skipping installation...")
		return
	}
	// Step 1: Get the latest Go version (e.g. "go1.20.5")
	versionCmd := exec.Command("sh", "-c", "curl -s https://go.dev/VERSION?m=text")
	versionOutput, err := versionCmd.Output()
	if err != nil {
		log.Fatal("Error fetching Go version:", err)
	}
	version := strings.TrimSpace(string(versionOutput))
	log.Println("Latest Go version:", version)

	// Step 2: Construct the download URL for Linux AMD64.
	tarballURL := fmt.Sprintf("https://go.dev/dl/%s.linux-amd64.tar.gz", version)
	log.Println("Downloading from:", tarballURL)

	// Step 3: Remove any previous Go installation, download and extract the new version.
	installCmdStr := fmt.Sprintf(
		"sudo rm -rf /usr/local/go && "+
			"sudo curl -L %s -o /tmp/go.tar.gz && "+
			"sudo tar -C /usr/local -xzf /tmp/go.tar.gz && "+
			"rm /tmp/go.tar.gz",
		tarballURL,
	)
	installCmd := exec.Command("sh", "-c", installCmdStr)
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	err = installCmd.Run()
	if err != nil {
		log.Fatal("Error installing Golang:", err)
	}

	// Step 4: Append Go environment variables to ~/.profile if they are not already set.
	// This snippet adds:
	//   # GO
	//   export PATH=$PATH:/usr/local/go/bin
	//   export PATH=$PATH:/usr/local/bin
	//   export PATH=~/go/bin:$PATH
	envCmdStr := "grep -q 'export PATH=.*/usr/local/go/bin' ~/.profile || " +
		"echo '\n# GO\nexport PATH=$PATH:/usr/local/go/bin\nexport PATH=$PATH:/usr/local/bin\n" +
		"export PATH=~/go/bin:$PATH' >> ~/.profile"
	envCmd := exec.Command("sh", "-c", envCmdStr)
	envCmd.Stdout = os.Stdout
	envCmd.Stderr = os.Stderr
	err = envCmd.Run()
	if err != nil {
		log.Fatal("Error updating ~/.profile:", err)
	}

	log.Println("Golang installed successfully and environment variables updated!")
	log.Println("")
}

func UninstallGolang() {
	log.Println("Uninstalling Golang...")

	// Step 1: Remove the Go installation directory.
	cmd := exec.Command("sh", "-c", "sudo rm -rf /usr/local/go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to remove /usr/local/go: ", err)
	}

	// Step 2 (Optional): Remove the user's Go workspace directory.
	cmd = exec.Command("sh", "-c", "rm -rf ~/go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println("Warning: Could not remove ~/go: ", err)
	}

	// Step 3: Clean up environment variables from ~/.profile.
	// This command removes the block starting with "# GO" and the next three lines.
	cmd = exec.Command("sh", "-c", "sed -i '/# GO/,+3d' ~/.profile")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to clean up ~/.profile: ", err)
	}

	log.Println("Golang has been uninstalled and .profile cleaned up!")
}


func UpgradeGolang() {
	log.Println("Checking for Golang upgrade...")

	// Step 1: Check the installed version by running "go version".
	versionCmd := exec.Command("sh", "-c", "go version")
	output, err := versionCmd.Output()
	if err != nil {
		log.Println("Golang is not currently installed; installing it now.")
		InstallGolang()
		return
	}

	fields := strings.Fields(string(output))
	if len(fields) < 3 {
		log.Println("Unexpected output from 'go version'; reinstalling Golang.")
		InstallGolang()
		return
	}

	installedVersion := fields[2] // e.g. "go1.24.2"
	log.Println("Currently installed Golang version:", installedVersion)

	// Step 2: Get the latest version from the official Go server.
	latestCmd := exec.Command("sh", "-c", "curl -s https://go.dev/VERSION?m=text")
	latestOutput, err := latestCmd.Output()
	if err != nil {
		log.Fatal("Error fetching latest Go version:", err)
	}
	// Again, use only the first line in case additional data is present.
	lines := strings.Split(string(latestOutput), "\n")
	latestVersion := strings.TrimSpace(lines[0])
	log.Println("Latest available Golang version:", latestVersion)

	// Step 3: Compare versions and upgrade if they differ.
	if installedVersion == latestVersion {
		log.Println("Golang is already up-to-date.")
		return
	}

	log.Printf("Upgrading Golang from %s to %s\n", installedVersion, latestVersion)
	InstallGolang()
}


func UpgradeTinygo() {
	log.Println("Checking for Tinygo upgrade...")

	// Step 1: Determine the currently installed Tinygo version.
	installedVersion := ""
	versionCmd := exec.Command("sh", "-c", "tinygo version")
	output, err := versionCmd.Output()
	if err != nil {
		log.Println("Tinygo is not installed.")
	} else {
		// Expected output: "tinygo version 0.37.0 (commit ...)"
		parts := strings.Fields(string(output))
		if len(parts) >= 3 {
			installedVersion = parts[2] // e.g., "0.37.0"
		}
		log.Println("Currently installed Tinygo version:", installedVersion)
	}

	// Step 2: Get the latest Tinygo version from the GitHub API.
	resp, err := http.Get("https://api.github.com/repos/tinygo-org/tinygo/releases/latest")
	if err != nil {
		log.Fatal("Error fetching latest Tinygo release:", err)
	}
	defer resp.Body.Close()

	type githubRelease struct {
		TagName string `json:"tag_name"`
	}
	var rel githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		log.Fatal("Error parsing Tinygo release JSON:", err)
	}

	latestTag := rel.TagName // e.g., "v0.37.0"
	latestVersion := strings.TrimPrefix(latestTag, "v")
	log.Println("Latest Tinygo version:", latestVersion)

	// Step 3: Compare versions. If the installed version matches latest, do nothing.
	if installedVersion == latestVersion {
		log.Println("Tinygo is already up-to-date.")
		return
	}

	log.Printf("Upgrading Tinygo from %s to %s\n", installedVersion, latestVersion)

	// Step 4: Construct the download URL and download the .deb package.
	debURL := fmt.Sprintf(
		"https://github.com/tinygo-org/tinygo/releases/download/v%s/tinygo_%s_amd64.deb",
		latestVersion, latestVersion,
	)
	log.Println("Downloading Tinygo from:", debURL)
	debFilePath := fmt.Sprintf("/tmp/tinygo_%s_amd64.deb", latestVersion)
	wgetCmdStr := fmt.Sprintf("wget -O %s %s", debFilePath, debURL)
	wgetCmd := exec.Command("sh", "-c", wgetCmdStr)
	wgetCmd.Stdout = os.Stdout
	wgetCmd.Stderr = os.Stderr
	if err := wgetCmd.Run(); err != nil {
		log.Fatalf("Error downloading Tinygo package: %v", err)
	}

	// Step 5: Install the downloaded .deb package using dpkg.
	dpkgCmdStr := fmt.Sprintf("sudo dpkg -i %s", debFilePath)
	dpkgCmd := exec.Command("sh", "-c", dpkgCmdStr)
	dpkgCmd.Stdout = os.Stdout
	dpkgCmd.Stderr = os.Stderr
	if err := dpkgCmd.Run(); err != nil {
		log.Fatalf("Error installing Tinygo: %v", err)
	}

	log.Println("Tinygo upgraded successfully!")
}


func InstallTinygo() {
	log.Println("Installing Tinygo...")

	// Step 1: Get the latest Tinygo release info from GitHub API.
	resp, err := http.Get("https://api.github.com/repos/tinygo-org/tinygo/releases/latest")
	if err != nil {
		log.Fatalf("Error fetching latest Tinygo release: %v", err)
	}
	defer resp.Body.Close()

	type githubRelease struct {
		TagName string `json:"tag_name"`
	}
	var rel githubRelease
	if err = json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		log.Fatalf("Error decoding Tinygo release JSON: %v", err)
	}

	// Extract the version, e.g. "v0.37.0" -> "0.37.0"
	latestTag := rel.TagName
	latestVersion := strings.TrimPrefix(latestTag, "v")
	log.Println("Latest Tinygo version:", latestVersion)

	// Step 2: Construct the download URL for the .deb package.
	debURL := fmt.Sprintf(
		"https://github.com/tinygo-org/tinygo/releases/download/v%s/tinygo_%s_amd64.deb",
		latestVersion, latestVersion,
	)
	log.Println("Downloading Tinygo from:", debURL)

	// Step 3: Download the package to /tmp.
	debFilePath := fmt.Sprintf("/tmp/tinygo_%s_amd64.deb", latestVersion)
	wgetCmdStr := fmt.Sprintf("wget -O %s %s", debFilePath, debURL)
	wgetCmd := exec.Command("sh", "-c", wgetCmdStr)
	wgetCmd.Stdout = os.Stdout
	wgetCmd.Stderr = os.Stderr
	if err = wgetCmd.Run(); err != nil {
		log.Fatalf("Error downloading Tinygo package: %v", err)
	}

	// Step 4: Install Tinygo using dpkg.
	dpkgCmdStr := fmt.Sprintf("sudo dpkg -i %s", debFilePath)
	dpkgCmd := exec.Command("sh", "-c", dpkgCmdStr)
	dpkgCmd.Stdout = os.Stdout
	dpkgCmd.Stderr = os.Stderr
	if err = dpkgCmd.Run(); err != nil {
		log.Fatalf("Error installing Tinygo: %v", err)
	}

	// Optionally, remove the downloaded .deb file after installation.
	if err = os.Remove(debFilePath); err != nil {
		log.Printf("Warning: Could not remove %s: %v", debFilePath, err)
	}

	log.Println("Tinygo installed successfully!")
}

func InstallGrowGD() {
	// go install graphics.gd/cmd/gd@master
	cmd := exec.Command("go", "install", "graphics.gd/cmd/gd@master")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to install GrowGD: ", err)
	}
}

func InstallCobraCli() {
	// go install github.com/spf13/cobra-cli@latest
	cmd := exec.Command("go", "install", "github.com/spf13/cobra-cli@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to install cobra-cli: ", err)
	}
}

func InstallGoose() {
	// go install github.com/pressly/goose/v3@latest
	cmd := exec.Command("go", "install", "github.com/pressly/goose/v3/cmd/goose@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to install goose: ", err)
	}
}