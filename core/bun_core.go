package core

import (
	"os"
	"os/exec"
	"github.com/charmbracelet/log"
)


func InstallBun() {
	// check if bun is already installed
	if _, err := exec.LookPath("bun"); err == nil {
		log.Info("bun is already installed.")
		return
	}

	// install bun
	log.Info("Installing bun")
	cmd := exec.Command("sh", "-c", "curl -fsSL https://bun.sh/install | bash")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func UninstallBun() {
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

func UpgradeBun() {
	log.Info("Updating bun")
	cmd := exec.Command("sh", "-c", "bun upgrade")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}