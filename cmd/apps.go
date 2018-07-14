// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/flogowebloader/util"
	"github.com/spf13/cobra"
)

// appsCmd represents the apps command
var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Apps management for the Project Flogo Web UI",
	Run:   runApps,
}

// Flags
var (
	flogoWebHost  string
	flogoFileName string
	readDir       bool
)

// Variables
var ()

const (
	// The API endpoint for Flogo Web UI
	flogoAPIEndpoint string = "/api/v2/"

	// The app import endpoint for Flogo Web UI
	flogoAppImport string = "apps:import"

	// The app endpoint for Flogo Web UI
	flogoApps string = "apps"
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(appsCmd)
}

// runPApps is the actual execution of the command
func runApps(cmd *cobra.Command, args []string) {
	fmt.Printf("\nThe Apps command supports the app management capabilities.\nThe commands available are:\n\n")

	// Print all subcommands
	for _, command := range cmd.Commands() {
		if command.Use != "help [command]" {
			fmt.Printf("%s %s\n", util.RightPadToLen(command.Use, ".", 25), command.Short)
		}
	}

	fmt.Printf("\nRun 'flogowebloader help apps [command]' for more details\n\n")
}
