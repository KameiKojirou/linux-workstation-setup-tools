/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"linux-workstation-setup-tools/core"

	"github.com/spf13/cobra"
)

// androidStudioCmd represents the androidStudio command
var androidStudioCmd = &cobra.Command{
	Use:   "androidStudio",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("androidStudio called")
		if  len (args) == 0 {
			fmt.Println("Please specify a command")
		}
		action, _ := cmd.Flags().GetString("action")
		switch action {
			case "install":
				core.InstallAndroidStudio()
			case "upgrade":
				core.UpgradeAndroidStudio()
			case "remove":
				core.UninstallAndroidStudio()
		}
	},
}

func init() {
	rootCmd.AddCommand(androidStudioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// androidStudioCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	androidStudioCmd.Flags().StringP("action", "a", "", "action to perform")

}
