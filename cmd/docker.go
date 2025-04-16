/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"linux-workstation-setup-tools/core"
	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker called")
		if  len (args) == 0 {
			fmt.Println("Please specify a command")
		}

		action, _ := cmd.Flags().GetString("action")
		switch action {
			case "install":
				core.InstallDocker()
			case "uninstall":
				core.UninstallDocker()
			default:
				fmt.Println("Please specify a valid action")
		}

		packages, _ := cmd.Flags().GetStringArray("packages")
		basepath, _ := cmd.Flags().GetString("basepath")
		if basepath == "" || len(packages) == 0 {
			fmt.Println("Please specify a basepath")
			return
		} else {
			for _, pkg := range packages {
				switch pkg {
					case "stirling-pdf":
						core.InstallStirlingPDFContainer(basepath)
					case "watchtower":
						core.InstallWatchtowerContainer(basepath)
					case "yacht":
						core.InstallYachtContainer(basepath)
					case "penpot":
						core.InstallPenPotContainer(basepath)
					default:
						fmt.Println("Please specify a valid package")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dockerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	dockerCmd.Flags().StringP("action", "a", "", "action to perform")
	dockerCmd.Flags().StringArrayP("packages", "p", []string{}, "list of packages to install")
	dockerCmd.Flags().StringP("basepath", "b", "", "base path for installation")
}
