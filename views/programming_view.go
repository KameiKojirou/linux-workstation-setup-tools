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
			huh.NewOption("Android-Studio", "android-studio"),
			huh.NewOption("Bun", "bun"),
			huh.NewOption("Deno", "deno"),
			huh.NewOption("Dotnet", "dotnet"),
			huh.NewOption("Go", "go"),
			huh.NewOption("Java", "java"),
			huh.NewOption("Rust", "rust"),
			huh.NewOption("UV", "uv"),
			huh.NewOption("Main Menu", "main"),
			huh.NewOption("Exit","exit"),
		).
		Value(&value)
		form.Run()


		switch value {
			case "android-studio":
				AndroidStudioMenu()
			case "bun":
				BunMenu()
			case "deno":
				DenoMenu()
			case "dotnet":
				DotnetCoreMenu()
			case "go":
				GoLangMenu()
			case "java":
				JavaMenu()
			case "rust":
				RustMenu()
			case "uv":
				UvMenu()
			case "main":
				MainMenu()
			case "exit":
				os.Exit(0)
		}
}