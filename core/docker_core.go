package core

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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



// InstallDocker and UninstallDocker functions remain unchanged

func InstallYachtContainer(basePath string) {
	// check if basePath is empty if so throw Error
	if basePath == "" {
		log.Fatal("basePath cannot be empty")
		return
	}

	log.Println("Installing Yacht container...")
	// Step 1: Check if the "yacht" container already exists.
	checkCmd := exec.Command("docker", "ps", "-a", "--filter", "name=^yacht$", "--format", "{{.Names}}")
	out, err := checkCmd.Output()
	if err != nil {
		log.Fatalf("Error checking for existing container: %v", err)
	}
	if containerName := strings.TrimSpace(string(out)); containerName != "" {
		log.Printf("Container '%s' is already installed. Skipping installation.\n", containerName)
		return
	}

	// Ensure config directory exists
	configPath := filepath.Join(basePath, "yacht", "config")
	if err := os.MkdirAll(configPath, 0755); err != nil {
		log.Fatalf("Error creating config directory: %v", err)
	}

	// Step 2: Construct the docker run command with proper flag formatting
	command := fmt.Sprintf(`docker run -d \
		--name yacht \
		--restart always \
		-p 8000:8000 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v "%s:/config" \
		selfhostedpro/yacht`, configPath)

	// Step 3: Execute the docker run command.
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error installing Yacht container: %v", err)
	}
	
	log.Println("Yacht container installed successfully!")
}

func InstallWatchtowerContainer(basePath string) {
	// check if basePath is empty if so throw Error
	if basePath == "" {
		log.Fatal("basePath cannot be empty")
		return
	}
	log.Println("Installing Watchtower container...")
	
	// Step 1: Check if the "watchtower" container already exists.
	checkCmd := exec.Command("docker", "ps", "-a", "--filter", "name=^watchtower$", "--format", "{{.Names}}")
	out, err := checkCmd.Output()
	if err != nil {
		log.Fatalf("Error checking for existing container: %v", err)
	}
	if containerName := strings.TrimSpace(string(out)); containerName != "" {
		log.Printf("Container '%s' is already installed. Skipping installation.\n", containerName)
		return
	}

	// Create config directory with sudo
	configPath := filepath.Join(basePath, "watchtower", "config")
	mkdirCmd := exec.Command("sudo", "mkdir", "-p", configPath)
	if err := mkdirCmd.Run(); err != nil {
		log.Fatalf("Error creating config directory: %v", err)
	}
	
	// Get current username for ownership change
	currentUser := os.Getenv("USER")
	if currentUser != "" {
		chownCmd := exec.Command("sudo", "chown", "-R", 
			fmt.Sprintf("%s:%s", currentUser, currentUser), 
			configPath)
		if err := chownCmd.Run(); err != nil {
			log.Printf("Warning: Could not change directory ownership: %v", err)
			// Continue anyway
		}
	}

	log.Printf("Using config path: %s\n", configPath)

	// Step 2: Construct the docker run command - Note: using the absolute path variable directly
	command := fmt.Sprintf(`docker run -d \
		--name watchtower \
		--restart unless-stopped \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v %s:/config \
		containrrr/watchtower`, configPath)

	// Step 3: Execute the docker run command.
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error installing Watchtower container: %v", err)
	}
	
	log.Println("Watchtower container installed successfully!")
}


func InstallStirlingPDFContainer(basePath string) {
	// check if basePath is empty if so throw Error
	if basePath == "" {
		log.Fatal("basePath cannot be empty")
		return
	}

	log.Println("Installing Stirling PDF container...")
	// Step 1: Check if the "stirling-pdf" container already exists.
	checkCmd := exec.Command("docker", "ps", "-a", "--filter", "name=^stirling-pdf$", "--format", "{{.Names}}")
	out, err := checkCmd.Output()
	if err != nil {
		log.Fatalf("Error checking for existing container: %v", err)
	}
	if containerName := strings.TrimSpace(string(out)); containerName != "" {
		log.Printf("Container '%s' is already installed. Skipping installation.\n", containerName)
		return
	}

	// Create all required directories
	dirs := []string{
		filepath.Join(basePath, "stirling-pdf", "trainingData"),
		filepath.Join(basePath, "stirling-pdf", "extraConfigs"),
		filepath.Join(basePath, "stirling-pdf", "customFiles"),
		filepath.Join(basePath, "stirling-pdf", "logs"),
		filepath.Join(basePath, "stirling-pdf", "pipeline"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatalf("Error creating directory %s: %v", dir, err)
		}
	}

	// Step 2: Construct the docker run command with absolute paths
	command := fmt.Sprintf(`docker run -d \
		--name stirling-pdf \
		--restart unless-stopped \
		-p 8080:8080 \
		-v "%s:/usr/share/tessdata" \
		-v "%s:/configs" \
		-v "%s:/customFiles" \
		-v "%s:/logs" \
		-v "%s:/pipeline" \
		-e DOCKER_ENABLE_SECURITY=false \
		-e LANGS=en_GB \
		docker.stirlingpdf.com/stirlingtools/stirling-pdf:latest`, 
		dirs[0], dirs[1], dirs[2], dirs[3], dirs[4])

	log.Println("Running docker command:")
	log.Println(command)

	// Step 3: Execute the docker run command.
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to run Docker container: ", err)
	}

	log.Println("Stirling PDF container deployed successfully!")
}

func InstallHomarrContainer(basePath string) {
	// check if basePath is empty if so throw Error
	if basePath == "" {
		log.Fatal("basePath cannot be empty")
		return
	}
	
	log.Println("Installing Homarr container...")
	
	// Step 1: Check if the "homarr" container already exists.
	checkCmd := exec.Command("docker", "ps", "-a", "--filter", "name=^homarr$", "--format", "{{.Names}}")
	out, err := checkCmd.Output()
	if err != nil {
		log.Fatalf("Error checking for existing container: %v", err)
	}
	if containerName := strings.TrimSpace(string(out)); containerName != "" {
		log.Printf("Container '%s' is already installed. Skipping installation.\n", containerName)
		return
	}

	// Create appdata directory
	appdataPath := filepath.Join(basePath, "homarr", "appdata")
	if err := os.MkdirAll(appdataPath, 0755); err != nil {
		log.Fatalf("Error creating appdata directory: %v", err)
	}

	randomBytes := make([]byte, 32) // 32 bytes = 64 hex characters
	if _, err := rand.Read(randomBytes); err != nil {
		log.Fatalf("Error generating random encryption key: %v", err)
	}
	encryptionKey := hex.EncodeToString(randomBytes)
	os.WriteFile("./homarrencryption.key", []byte(encryptionKey), 0600)

	// Step 2: Construct the docker run command with proper paths
	command := fmt.Sprintf(`docker run -d \
		--name homarr \
		--restart unless-stopped \
		-p 80:7575 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v "%s:/appdata" \
		-e SECRET_ENCRYPTION_KEY='%s' \
		ghcr.io/homarr-labs/homarr:latest`, appdataPath, encryptionKey)

	// Step 3: Execute the docker run command.
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to run Docker container: ", err)
	}
	
	log.Println("Homarr container installed successfully!")
}

func InstallPenPotContainer(basePath string) {
	// TODO implementation
	log.Println("PenPot container installation not implemented yet!")
}