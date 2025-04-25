package views

import (
	"fmt"
	"linux-workstation-setup-tools/core"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"slices"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

func ContainersMenu() {
	var value string
	form := huh.NewSelect[string]().
		Title("Containers Menu").
		Options(
			huh.NewOption("Manage Containers", "management-containers"),
			huh.NewOption("Main Menu", "main"),
			huh.NewOption("Exit", "exit"),
		).
		Value(&value)
	form.Run()

	switch value {
	case "management-containers":
		ManagementContainersMenu()
	case "main":
		MainMenu()
	case "exit":
		os.Exit(0)
	}
}

// Function to check if a container is running
func isContainerRunning(containerName string) bool {
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{.Names}}")
	output, err := cmd.Output()
	if err != nil {
		log.Error("Failed to check container status", "error", err)
		return false
	}

	containers := strings.Split(string(output), "\n")
	for _, container := range containers {
		if strings.TrimSpace(container) == containerName {
			return true
		}
	}
	return false
}

func ManagementContainersMenu() {
	// Define the containers we want to manage
	containers := []string{"yacht", "watchtower", "stirling-pdf", "homarr"}
	
	// Check current state of containers
	initialState := make(map[string]bool)
	for _, container := range containers {
		initialState[container] = isContainerRunning(container)
	}
	
	// Create options with pre-selected state based on what's already installed
	var selectedContainers []string
	options := make([]huh.Option[string], 0, len(containers))
	
	for _, container := range containers {
		options = append(options, huh.NewOption(capitalize(container), container))
		if initialState[container] {
			selectedContainers = append(selectedContainers, container)
		}
	}
	
	// Create the form with pre-selected values
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Manage Docker Containers").
				Description("Select containers to install, deselect to uninstall").
				Options(options...).
				Value(&selectedContainers),
		),
	)
	
	// Run the form
	err := form.Run()
	if err != nil {
		log.Error("Error running form", "error", err)
		return
	}
	
	// Process changes based on selection differences
	for _, container := range containers {
		wasSelected := initialState[container]
		isNowSelected := contains(selectedContainers, container)
		
		if !wasSelected && isNowSelected {
			// Container was added - install it
			log.Info("Installing container", "container", container)
			installContainer(container)
		} else if wasSelected && !isNowSelected {
			// Container was removed - uninstall it
			log.Info("Uninstalling container", "container", container)
			uninstallContainer(container)
		}
	}
	
	// Return to containers menu
	ContainersMenu()
}

// Helper function to install the appropriate container
// All other functions remain the same...

// Helper function to install the appropriate container
// Helper function to install the appropriate container
func installContainer(container string) {
	// Get the current user's home directory
	currentUser, err := user.Current()
	if err != nil {
		log.Error("Failed to get current user", "error", err)
		return
	}
	
	// Create absolute path to ~/containers
	containerPath := filepath.Join(currentUser.HomeDir, "containers")
	
	// Ensure the container directory exists using sudo if needed
	mkdirCmd := exec.Command("sudo", "mkdir", "-p", containerPath)
	if err := mkdirCmd.Run(); err != nil {
		log.Error("Failed to create containers directory", "error", err)
		return
	}
	
	// Change ownership to current user
	chownCmd := exec.Command("sudo", "chown", "-R", 
		fmt.Sprintf("%s:%s", currentUser.Username, currentUser.Username), 
		containerPath)
	if err := chownCmd.Run(); err != nil {
		log.Error("Failed to change directory ownership", "error", err)
		// Continue anyway as the directory might still be usable
	}
	
	log.Info("Installing to path", "path", containerPath)
	
	switch container {
	case "yacht":
		core.InstallYachtContainer(containerPath)
	case "watchtower":
		core.InstallWatchtowerContainer(containerPath)
	case "stirling-pdf":
		core.InstallStirlingPDFContainer(containerPath)
	case "homarr":
		core.InstallHomarrContainer(containerPath)
	default:
		log.Warn("Unknown container type", "container", container)
	}
}


// Helper function to uninstall a container
// ... rest of the functions remain unchanged

// Helper function to uninstall a container
func uninstallContainer(container string) {
	// Stop and remove the container
	stopCmd := exec.Command("docker", "stop", container)
	if err := stopCmd.Run(); err != nil {
		log.Error("Failed to stop container", "container", container, "error", err)
	}
	
	rmCmd := exec.Command("docker", "rm", container)
	if err := rmCmd.Run(); err != nil {
		log.Error("Failed to remove container", "container", container, "error", err)
	}
	
	log.Info("Container uninstalled successfully", "container", container)
}

// Helper function to check if a string is in a slice
func contains(slice []string, item string) bool {
	return slices.Contains(slice, item)
}

// Helper function to capitalize first letter
func capitalize(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}