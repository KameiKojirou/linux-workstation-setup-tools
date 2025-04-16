/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"linux-workstation-setup-tools/core"

	"github.com/spf13/cobra"
)

// uvCmd represents the uv command
var uvCmd = &cobra.Command{
	Use:   "uv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uv called")
		if  len (args) == 0 {
			fmt.Println("Please specify a command")
		}
		action, _ := cmd.Flags().GetString("action")
		switch action {
			case "install":
				core.InstallUv()
			case "update":
				core.UpdateUv()
			case "uninstall":
				core.UninstallUv()
		}
	},
}

func init() {
	rootCmd.AddCommand(uvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	uvCmd.Flags().StringP("action", "a", "", "action to perform")
}
