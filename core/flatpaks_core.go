package core

import (
	"os"
	"os/exec"
	"github.com/charmbracelet/log"
)


/*
 ---FLATPAKS---
*/


/* Firefox */

func InstallFlatpakFirefox() {
	// flatpak install flathub org.mozilla.firefox
	cmd := exec.Command("flatpak", "install", "flathub", "org.mozilla.firefox")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to install Flatpak Firefox: ", err)
	}
}

func UninstallFlatpakFirefox() {
	// flatpak uninstall org.mozilla.firefox
	cmd := exec.Command("flatpak", "uninstall", "org.mozilla.firefox")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to uninstall Flatpak Firefox: ", err)
	}
	
}