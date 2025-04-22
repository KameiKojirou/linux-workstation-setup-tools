package views

import (
	"linux-workstation-setup-tools/core"
	"linux-workstation-setup-tools/components"
	"os"
	"github.com/charmbracelet/huh"
)

func DotnetCoreMenu() {
	value := ""
	form := huh.NewSelect[string]().
		Title("Dotnet Core Menu").
		Options(
			huh.NewOption("Install", "install"),
			huh.NewOption("Update", "update"),
			huh.NewOption("Remove", "remove"),
			huh.NewOption("Main Menu", "main"),
			huh.NewOption("Exit","exit"),
		).
		Value(&value)
		form.Run()

		components.ConfirmationCheck()

		switch value {
			case "install":
				core.InstallDotnetCore()
			case "upgrade":
				core.UpdateDotnetCore()
			case "remove":
				core.UninstallDotnetCore()
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}
}