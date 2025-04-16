package core

import (
	"os"
	"os/exec"
	"github.com/charmbracelet/log"
)


func InstallBunJs() {
	log.Info("Installing bun")
	cmd := exec.Command("sh", "-c", "curl -fsSL https://bun.sh/install | bash")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func UninstallBunJs() {
	// remove bunjs from install and path
	log.Info("Uninstalling bun")
	cmd := exec.Command("sh", "-c", "bun uninstall")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateBunJs() {
	log.Info("Updating bun")
	cmd := exec.Command("sh", "-c", "bun upgrade")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}