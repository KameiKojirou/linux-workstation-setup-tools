package core

import (
	"os"
	"os/exec"
	"github.com/charmbracelet/log"
)

func InstallUv() {
	// curl -LsSf https://astral.sh/uv/install.sh | sh
	log.Info("Installing astralUv")
	cmd := exec.Command("sh", "-c", "curl -LsSf https://astral.sh/uv/install.sh | sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func UninstallUv() {
	log.Info("Uninstalling astralUv")
	cmd := exec.Command("sh", "-c", "astral uninstall")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}


func UpdateUv() {
	log.Info("Updating astralUv")
	cmd := exec.Command("uv", "self", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}