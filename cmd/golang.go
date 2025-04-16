/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"linux-workstation-setup-tools/core"
	"github.com/spf13/cobra"
)

// golangCmd represents the golang command
var golangCmd = &cobra.Command{
	Use:   "golang",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if  len (args) == 0 {
			fmt.Println("Please specify a command")
		}
		action, _ := cmd.Flags().GetString("action")
		switch action {
			case "install":
				core.InstallGolang()
			case "upgrade":
				core.UpgradeGolang()
			case "uninstall":
				core.UninstallGolang()
			default:
				fmt.Println("Please specify a valid action")
		}
		programs, _ := cmd.Flags().GetStringArray("programs")
		for _, program := range programs {
			switch program {
				case "tinygo":
					core.InstallTinygo()
				case "grow":
					core.InstallGrowGD()
				case "goose":
					core.InstallGoose()
				default:
					fmt.Println("Please specify a valid program")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(golangCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// golangCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	golangCmd.Flags().StringP("action", "a", "", "action to perform")
	golangCmd.Flags().StringArrayP("programs", "p", []string{}, "list of programs to install")
}
