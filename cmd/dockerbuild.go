// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/retgits/flogowebloader/util"
	"github.com/spf13/cobra"
)

// dockerBuildCmd represents the dockerBuild command
var dockerBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds a new docker image",
	Run:   runDockerBuild,
}

// Flags
var ()

// Variables
var ()

// init registers the command and flags
func init() {
	dockerCmd.AddCommand(dockerBuildCmd)
	dockerBuildCmd.Flags().StringVar(&imageName, "image", "flogo/flogo-docker", "The image name for the Flogo Web UI container")
	dockerBuildCmd.Flags().StringVar(&importsFile, "imports", "", "An imports file in case you want to add additional activities (like /home/user/Downloads/imports.go)")
}

// dockerfile is the template for a dockerfile needed to build a docker image
const dockerfile = `# This Dockerfile allows you to build a Flogo web docker image which comes
# preinstalled with additional activities. To add activities update the
# imports.go file and for each activity you want to have installed add the
# go get path (which you also use for flogo install). So if the URL you would
# paste is http://github.com/retgits/flogo-components/activity/trellocard
# simply add _ "github.com/retgits/flogo-components/activity/trellocard" to the
# imports.go
#
# Build a new image
# docker build . -t {{.name}}
#
# Run the image:
# docker run -d -p 3303:3303 --name=flogoweb {{.name}} eula-accept
#
FROM flogo/flogo-docker:latest

# Add the imports.go which contains all activities you want to pre-install and
# run flogo ensure to make sure the activities are installed
ADD imports.go /tmp/flogo-web/build/server/local/engines/flogo-web/src/flogo-web/imports.go
RUN cd /tmp/flogo-web/build/server/local/engines/flogo-web && flogo ensure`

// runDockerBuild is the actual execution of the command
func runDockerBuild(cmd *cobra.Command, args []string) {
	// Create a temporary directory
	dir, err := ioutil.TempDir(util.CurrentDirectory(), "docker")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	defer os.RemoveAll(dir) // clean up

	// Prepare the dockerfile
	data := make(map[string]interface{})
	data["name"] = imageName

	t := template.Must(template.New("email").Parse(dockerfile))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	// Write the dockerfile to disk
	tmpfn := filepath.Join(dir, "dockerfile")
	if err := ioutil.WriteFile(tmpfn, buf.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}

	// Copy the imports
	err = util.CopyFile(importsFile, filepath.Join(dir, "imports.go"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	// Execute docker build
	cmdExec := exec.Command("docker", "build", ".", "-t", imageName)
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr
	cmdExec.Dir = dir
	cmdExec.Run()
}
