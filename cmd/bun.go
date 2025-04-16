/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"linux-workstation-setup-tools/core"
	"github.com/spf13/cobra"
)

// bunCmd represents the bun command
var bunCmd = &cobra.Command{
	Use:   "bun",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len (args) == 0 {
			fmt.Println("bun called")
			fmt.Println("Please specify a command")
		}
		action, _ := cmd.Flags().GetString("action")
		switch action {
			case "install":
				core.InstallBunJs()
			case "upgrade":
				core.UpdateBunJs()
			case "uninstall":
				core.UninstallBunJs()
		}
	},
}

func init() {
	rootCmd.AddCommand(bunCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	bunCmd.Flags().StringP("action", "a", "", "action to perform")
}
