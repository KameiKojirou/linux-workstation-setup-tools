package views

import (
	"linux-workstation-setup-tools/core"
	"linux-workstation-setup-tools/components"
	"os"
	"github.com/charmbracelet/huh"
)

func RustMenu() {
	value := ""
	form := huh.NewSelect[string]().
		Title("Rust Menu").
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
				core.InstallRust()
			case "upgrade":
				core.UpgradeRust()
			case "remove":
				core.UninstallRust()
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}
}