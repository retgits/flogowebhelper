// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// dockerStartCmd represents the dockerStart command
var dockerStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a new instance of the Flogo Web UI with default settings",
	Run:   runDockerStart,
}

// Flags
var ()

// Variables
var ()

// init registers the command and flags
func init() {
	dockerCmd.AddCommand(dockerStartCmd)
	dockerStartCmd.Flags().StringVar(&imageName, "image", "flogo/flogo-docker", "The image name for the Flogo Web UI container")
}

// runDockerStart is the actual execution of the command
func runDockerStart(cmd *cobra.Command, args []string) {
	cmdExec := exec.Command("docker", "inspect", "--format", "{{.State.Running}}", "flogoweb")

	output, err := cmdExec.CombinedOutput()
	if err != nil && !strings.Contains(string(output), "No such object") {
		fmt.Println(err.Error())
	}

	if strings.HasPrefix(string(output), "true") {
		fmt.Printf("\nFlogoweb is already running!\n\n")
		cmdExec = exec.Command("docker", "ps", "-f", "name=flogoweb")
		cmdExec.Stdout = os.Stdout
		cmdExec.Stderr = os.Stderr
		cmdExec.Run()
		fmt.Println()
		os.Exit(0)
	} else {
		cmdExec = exec.Command("docker", "rm", "flogoweb")
		cmdExec.Run()
		cmdExec = exec.Command("docker", "run", "-d", "--name=flogoweb", "-p", "3303:3303", imageName, "eula-accept")
		cmdExec.Stdout = os.Stdout
		cmdExec.Stderr = os.Stderr
		cmdExec.Run()
		fmt.Println()
	}
}
