package views

import (
	"fmt"
	"os"
	"github.com/charmbracelet/huh"
)


func ProgrammingMenu() {
	value := ""
	form := huh.NewSelect[string]().
		Title("Programming Languages").
		Options(
			huh.NewOption("Go", "go"),
			huh.NewOption("Rust", "rust"),
			huh.NewOption("Deno", "deno"),
			huh.NewOption("Main Menu", "main"),
			huh.NewOption("Exit","exit"),
		).
		Value(&value)
		form.Run()


		switch value {
			case "go":
				GoLangMenu()
			case "rust":
				fmt.Println("Rust")
			case "deno":
				fmt.Println("Deno")
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}
}