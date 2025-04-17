package core


import (
	"os"
	"os/exec"
	"github.com/charmbracelet/log"
)


func InstallRust() {
	// check if  rust is already installed
	if _, err := exec.LookPath("rustc"); err == nil {
		log.Info("rust is already installed.")
		return
	}
	log.Info("Installing rust")
	cmd := exec.Command("sh", "-c", "curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func UpgradeRust() {
	log.Info("Updating rust")
	cmd := exec.Command("sh", "-c", "rustup upgrade")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func UninstallRust() {
	log.Info("Uninstalling rust")
	cmd := exec.Command("sh", "-c", "rustup self uninstall -y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}