package views

import (
	"linux-workstation-setup-tools/components"
	"linux-workstation-setup-tools/core"
	"os"

	"github.com/charmbracelet/huh"
)

func DockerMenu() {
	var value string
	form := huh.NewSelect[string]().
		Title("Docker Menu").
		Options(
			huh.NewOption("Install", "install"),
			huh.NewOption("Uninstall", "uninstall"),
			huh.NewOption("Containers", "containers"),
			huh.NewOption("Main Menu", "main"),
			huh.NewOption("Exit", "exit"),
		).
		Value(&value)
	form.Run()

	switch value {
	case "install":
		components.ConfirmationCheck()
		core.InstallDocker()
	case "containers":
		ContainersMenu()
	case "uninstall":
		components.ConfirmationCheck()
		core.UninstallDocker()
	case "main":
		MainMenu()
	case "exit":
		os.Exit(0)
	}
}
