// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// dockerLatestCmd represents the dockerLatest command
var dockerLatestCmd = &cobra.Command{
	Use:   "latest",
	Short: "Pulls the latest version of the Flogo Web UI from Docker Hub",
	Run:   runDockerLatest,
}

// Flags
var ()

// Variables
var ()

// init registers the command and flags
func init() {
	dockerCmd.AddCommand(dockerLatestCmd)
}

// runDockerLatest is the actual execution of the command
func runDockerLatest(cmd *cobra.Command, args []string) {
	cmdExec := exec.Command("docker", "pull", "flogo/flogo-docker:latest")
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	cmdExec.Run()
}
