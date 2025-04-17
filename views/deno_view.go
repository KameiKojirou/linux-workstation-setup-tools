package views

import (
	"linux-workstation-setup-tools/components"
	"linux-workstation-setup-tools/core"
	"os"
	"github.com/charmbracelet/huh"
)


func DenoMenu() {
	value := ""
	form := huh.NewSelect[string]().
		Title("Deno Menu").
		Options(
			huh.NewOption("Install", "install"),
			huh.NewOption("Upgrade", "upgrade"),
			huh.NewOption("Remove", "remove"),
			huh.NewOption("Main Menu", "main"),
			huh.NewOption("Exit","exit"),
		).
		Value(&value)
		form.Run()

		components.ConfirmationCheck()

		switch value {
			case "install":
				core.InstallDeno()
			case "upgrade":
				core.UpgradeDeno()
			case "remove":
				core.UninstallDeno()
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}
}