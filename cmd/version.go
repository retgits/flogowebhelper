// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the current version",
	Run:   runVersion,
}

const version = "0.0.1"

// init registers the command and flags
func init() {
	rootCmd.AddCommand(versionCmd)
}

// runVersion is the actual execution of the command
func runVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("\nThe current version of the app is: %s\n\n", version)
}
