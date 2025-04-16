package core

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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



func  InstallYachtContainer(basePath string) {
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

	// Step 2: Construct the docker run command with dynamic volume mappings.
	/* docker volume create yacht
docker run -d -p 8000:8000 -v /var/run/docker.sock:/var/run/docker.sock -v yacht:/config selfhostedpro/yacht */

	command := fmt.Sprintf(`docker run -d \
		--name yacht \
		-p 8000:8000 \
		-restart always \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v "%s/yacht/config:/config" \
		selfhostedpro/yacht`, basePath)

	// Step 3: Execute the docker run command.
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error installing Yacht container: %v", err)
	}
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
	/*
	$ docker run --detach \
    --name watchtower \
    --volume /var/run/docker.sock:/var/run/docker.sock \
    containrrr/watchtower
	*/
	// Step 2: Construct the docker run command with dynamic volume mappings.

	command := fmt.Sprintf(`docker run -d \
		--name watchtower \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v "%s/watchtower/config:/config" \
		containrrr/watchtower`, basePath)

	// Step 3: Execute the docker run command.
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error installing Watchtower container: %v", err)
	}
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

	// Step 2: Construct the docker run command with dynamic volume mappings.
	command := fmt.Sprintf(`docker run -d \
		--name stirling-pdf \
		-p 8080:8080 \
		-v "%s/stirling-pdf/trainingData:/usr/share/tessdata" \
		-v "%s/stirling-pdf/extraConfigs:/configs" \
		-v "%s/stirling-pdf/customFiles:/customFiles" \
		-v "%s/stirling-pdf/logs:/logs" \
		-v "%s/stirling-pdf/pipeline:/pipeline" \
		-e DOCKER_ENABLE_SECURITY=false \
		-e LANGS=en_GB \
		docker.stirlingpdf.com/stirlingtools/stirling-pdf:latest`, basePath, basePath, basePath, basePath, basePath)

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


func  InstallPenPotContainer(basePath string) {
	// TODO!
}