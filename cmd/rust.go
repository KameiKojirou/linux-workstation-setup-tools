/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"linux-workstation-setup-tools/core"
	// "linux-workstation-setup-tools/core"
	"github.com/spf13/cobra"
)

// rustCmd represents the rust command
var rustCmd = &cobra.Command{
	Use:   "rust",
	Short: "install rust",
	Long: `install rust on your linux workstation`,
	Run: func(cmd *cobra.Command, args []string) {
		if len (args) == 0 {
			fmt.Println("Please specify a command")
		}
		action, _ := cmd.Flags().GetString("action")
		switch action {
			case "install":
				core.InstallRust()
			case "upgrade":
				core.UpgradeRust()
			case "uninstall":
				core.UninstallRust()
		}
	},
}

func init() {
	rootCmd.AddCommand(rustCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rustCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	rustCmd.Flags().StringP("action", "a", "", "action to perform")
}
