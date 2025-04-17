package views

import (
	"linux-workstation-setup-tools/core"
	"github.com/charmbracelet/huh"
)


func MainMenu() {
	value := ""
	form := huh.NewSelect[string]().
		Title("Main Menu").
		Options(
			huh.NewOption("Programming Languages", "programming"),
			huh.NewOption("Exit","exit"),
		).
		Value(&value)
		form.Run()

		switch value {
			case "programming":
				ProgrammingMenu()
			default:
				core.Exit()
		}
}