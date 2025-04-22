package views

import (
	"linux-workstation-setup-tools/core"
	"linux-workstation-setup-tools/components"
	"os"
	"github.com/charmbracelet/huh"
)

func UvMenu() {
	value := ""
	form := huh.NewSelect[string]().
		Title("UV Menu").
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
				core.InstallUV()
			case "upgrade":
				core.UpdateUV()
			case "remove":
				core.UninstallUV()
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}
}