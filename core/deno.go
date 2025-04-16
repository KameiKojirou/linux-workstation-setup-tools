package core

import (
	"os"
	"os/exec"
	"github.com/charmbracelet/log"
)


func InstallDeno() {
	log.Info("Installing deno")
	cmd := exec.Command("sh", "-c", "curl -fsSL https://deno.land/x/install/install.sh | sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}


func UninstallDeno() {
	log.Info("Uninstalling deno")
	cmd := exec.Command("sh", "-c", "deno uninstall")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}


func UpdateDeno() {
	log.Info("Updating deno")
	cmd := exec.Command("sh", "-c", "deno upgrade")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}