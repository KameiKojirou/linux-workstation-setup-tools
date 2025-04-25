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
			huh.NewOption("Tools", "tools"),
			huh.NewOption("Exit","exit"),
		).
		Value(&value)
		form.Run()

		switch value {
			case "programming":
				ProgrammingMenu()
			case "tools":
				ToolsView()
			default:
				core.Exit()
		}
}