// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/flogowebloader/util"
	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker container management for the Project Flogo Web UI",
	Run:   runDocker,
}

// Flags
var (
	imageName   string
	importsFile string
)

// Variables
var ()

// init registers the command and flags
func init() {
	rootCmd.AddCommand(dockerCmd)
}

// runDocker is the actual execution of the command
func runDocker(cmd *cobra.Command, args []string) {
	fmt.Printf("\nThe Docker command supports the container management capabilities.\nThe commands available are:\n\n")

	// Print all subcommands
	for _, command := range cmd.Commands() {
		if command.Use != "help [command]" {
			fmt.Printf("%s %s\n", util.RightPadToLen(command.Use, ".", 25), command.Short)
		}
	}

	fmt.Printf("\nRun 'flogowebloader help docker [command]' for more details\n\n")
}
