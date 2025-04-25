package views

import (
	"linux-workstation-setup-tools/core"
	"linux-workstation-setup-tools/components"
	"os"
	"github.com/charmbracelet/huh"
)

func TabbyMenu() {
	value := ""
	form := huh.NewSelect[string]().
		Title("Tabby Menu").
		Options(
			huh.NewOption("Install", "install"),
			huh.NewOption("Upgrade", "upgrade"),
			huh.NewOption("Uninstall", "uninstall"),
			huh.NewOption("Main Menu", "main"),
			huh.NewOption("Exit","exit"),
		).Value(&value)
		form.Run()

		components.ConfirmationCheck()

		switch value {
			case "install":
				core.InstallTabby()
			case "uninstall":
				core.UninstallTabby()
			case "upgrade":
				core.UpgradeTabby()
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}
} 