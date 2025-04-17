package views

import (
	"os"
	"github.com/charmbracelet/huh"
)


func ProgrammingMenu() {
	value := ""
	form := huh.NewSelect[string]().
		Title("Programming Languages").
		Options(
			huh.NewOption("Bun", "bun"),
			huh.NewOption("Deno", "deno"),
			huh.NewOption("Go", "go"),
			huh.NewOption("Rust", "rust"),
			huh.NewOption("Main Menu", "main"),
			huh.NewOption("Exit","exit"),
		).
		Value(&value)
		form.Run()


		switch value {
			case "bun":
				BunMenu()
			case "go":
				GoLangMenu()
			case "rust":
				RustMenu()
			case "deno":
				DenoMenu()
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}
}