package views

import (
	"linux-workstation-setup-tools/core"
	"linux-workstation-setup-tools/components"
	"os"
	"github.com/charmbracelet/huh"
)

func BunMenu() {
	value := ""
	form := huh.NewSelect[string]().
		Title("Bun Menu").
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
				core.InstallBun()
			case "upgrade":
				core.UpgradeBun()
			case "remove":
				core.UninstallBun()
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}
}