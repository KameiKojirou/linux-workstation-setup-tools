package core

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// InstallDocker installs Docker on Ubuntu (and Pop!_OS 22.04) following the
// official installation script. It will remove any existing Docker installation,
// install Docker, and add the current user to the Docker group.
func InstallDocker() {
	log.Println("Installing Docker...")

	// Step 1: Check if Docker is already installed.
	checkCmd := exec.Command("sh", "-c", "docker --version")
	checkCmd.Stdout = os.Stdout
	checkCmd.Stderr = os.Stderr
	if err := checkCmd.Run(); err == nil {
		log.Println("Docker is already installed. Removing current installation...")

		// Remove Docker packages and cleanup Docker directories.
		removeCmdStr := `
			sudo apt-get remove --purge -y docker-ce docker-ce-cli containerd.io docker docker-engine docker.io && \
			sudo rm -rf /var/lib/docker && \
			sudo rm -rf /var/lib/containerd
		`
		removeCmd := exec.Command("sh", "-c", removeCmdStr)
		removeCmd.Stdout = os.Stdout
		removeCmd.Stderr = os.Stderr
		if err := removeCmd.Run(); err != nil {
			log.Fatalf("Error removing current Docker installation: %v", err)
		}
	} else {
		log.Println("Docker not found. Proceeding with installation...")
	}

	// Step 2: Download the Docker installation script.
	curlCmd := exec.Command("sh", "-c", "curl -fsSL https://get.docker.com -o get-docker.sh")
	curlCmd.Stdout = os.Stdout
	curlCmd.Stderr = os.Stderr
	if err := curlCmd.Run(); err != nil {
		log.Fatalf("Error downloading Docker installation script: %v", err)
	}

	// Step 3: Run the Docker installation script.
	installCmd := exec.Command("sh", "-c", "sudo sh get-docker.sh")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	if err := installCmd.Run(); err != nil {
		log.Fatalf("Error installing Docker: %v", err)
	}

	// Step 4: Remove the downloaded installation script.
	if err := os.Remove("get-docker.sh"); err != nil {
		log.Printf("Warning: Could not remove get-docker.sh: %v", err)
	}

	// Step 5: Add the current user to the Docker group.
	currentUser := os.Getenv("USER")
	if currentUser == "" {
		log.Println("Error: USER environment variable is not set. Skipping adding user to docker group.")
	} else {
		groupCmdStr := fmt.Sprintf("sudo usermod -aG docker %s", currentUser)
		groupCmd := exec.Command("sh", "-c", groupCmdStr)
		groupCmd.Stdout = os.Stdout
		groupCmd.Stderr = os.Stderr
		if err := groupCmd.Run(); err != nil {
			log.Fatalf("Error adding user to docker group: %v", err)
		}
	}

	log.Println("Docker installed successfully!")
}



// UninstallDocker completely removes Docker and its related files.
func UninstallDocker() {
	log.Println("Uninstalling Docker...")

	// Remove Docker packages and cleanup Docker directories.
	removeCmdStr := `
sudo apt-get remove --purge -y docker-ce docker-ce-cli containerd.io docker docker-engine docker.io &&
sudo rm -rf /var/lib/docker &&
sudo rm -rf /var/lib/containerd
`
	cmd := exec.Command("sh", "-c", removeCmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error uninstalling Docker: %v", err)
	}

	log.Println("Docker has been uninstalled successfully!")
}
