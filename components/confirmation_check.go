package components

import (
	"github.com/charmbracelet/huh"
	"linux-workstation-setup-tools/core"
)

func ConfirmationCheck() {
	value := ""
	form := huh.NewSelect[string]().
		Title("Are you sure?").
		Options(
			huh.NewOption("Yes", "yes"),
			huh.NewOption("No", "no"),
		).
		Value(&value)
		form.Run()


		switch value {
			case "no":
				core.Exit()
		}
}