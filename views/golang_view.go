package views

import (
	"linux-workstation-setup-tools/components"
	"linux-workstation-setup-tools/core"
	"os"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
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

		components.ConfirmationCheck()

		switch value {
			case "install":
				core.InstallGolang()
				InstallGoLibraries()
			case "upgrade":
				core.UpgradeGolang()
			case "remove":
				core.UninstallGolang()
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}

		GoLangMenu()
	}

func InstallGoLibraries() {
	var multivalue []string
	librariesform := huh.NewMultiSelect[string]().
		Title("Golang Libraries").
		Options(
			huh.NewOption("Tinygo", "tinygo"),
			huh.NewOption("Grow", "grow"),
			huh.NewOption("Cobra-Cli", "cobra"),
		).
		Value(&multivalue)
		librariesform.Run()

		fmt.Println("you selected: ",multivalue)

		components.ConfirmationCheck()

		if  len(multivalue) > 1 {
			for _, v := range multivalue {
				switch v {
					case "tinygo":
						core.InstallTinygo()
					case "grow":
						core.InstallGrowGD()
					case "cobra":
						core.InstallCobraCli()
				}
			}

		} else {
			log.Info("No libraries selected")
			GoLangMenu()
		}

}