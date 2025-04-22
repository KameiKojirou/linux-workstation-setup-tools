package core

import (
	"os"
	"os/exec"
	"github.com/charmbracelet/log"
)

func InstallUV() {
	// check  if  astralUv is already installed
	if _, err := exec.LookPath("uv"); err == nil {
		log.Info("astralUv is already installed.")
		return
	}
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

func UninstallUV() {
	log.Info("Uninstalling astralUv")
	cmd := exec.Command("sh", "-c", "astral uninstall")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}


func UpdateUV() {
	log.Info("Updating astralUv")
	cmd := exec.Command("uv", "self", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}