package views

import (
	"linux-workstation-setup-tools/core"
	"os"
	"github.com/charmbracelet/huh"
)


func GoLangMenu() {
	value := ""
	form := huh.NewSelect[string]().
		Title("Golang Menu").
		Options(
			huh.NewOption("Install", "install"),
			huh.NewOption("Upgrade", "upgrade"),
			huh.NewOption("Remove", "remove"),
			huh.NewOption("Main Menu", "main"),
			huh.NewOption("Exit","exit"),
		).
		Value(&value)
		form.Run()


		switch value {
			case "install":
				core.InstallGolang()
			case "upgrade":
				core.UpgradeGolang()
			case "remove":
				core.UninstallGolang()
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}

}