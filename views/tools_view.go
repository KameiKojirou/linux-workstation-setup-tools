package views

import (
	"os"
	"github.com/charmbracelet/huh"
)

func ToolsView() {
	var value string
	form := huh.NewSelect[string]().
		Title("Tools Menu").
		Options(
			huh.NewOption("Docker", "docker"),
			huh.NewOption("Tabby", "tabby"),
			huh.NewOption("Main Menu", "main"),
			huh.NewOption("Exit","exit"),
		).
		Value(&value)
		form.Run()

		switch value {
			case "docker":
				DockerMenu()
			case "tabby":
				TabbyMenu()
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}
}

