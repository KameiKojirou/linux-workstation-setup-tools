/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"linux-workstation-setup-tools/core"

	"github.com/spf13/cobra"
)

// flatpaksCmd represents the flatpaks command
var flatpaksCmd = &cobra.Command{
	Use:   "flatpaks",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("flatpaks called")
		if  len (args) == 0 {
			fmt.Println("Please specify a command")
			return
		}
		flatpaks, _ := cmd.Flags().GetStringArray("flatpaks")
		if len(flatpaks) == 0 {
			fmt.Println("Please specify a flatpak")
			return
		}

		for _, flatpak := range flatpaks {
			switch flatpak {
				case "firefox":
					core.InstallFlatpakFirefox()
				default:
					fmt.Println("Please specify a valid flatpak")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(flatpaksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flatpaksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	flatpaksCmd.Flags().StringArrayP("flatpaks", "f", []string{}, "list of flatpaks to install")
}
