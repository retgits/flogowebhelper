// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/retgits/flogowebhelper/util"
	"github.com/spf13/cobra"
)

// appsImportCmd represents the appsImport command
var appsImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports Flogo apps into the Flogo Web UI",
	Run:   runAppsImport,
}

// Flags
var ()

// Variables
var ()

// init registers the command and flags
func init() {
	appsCmd.AddCommand(appsImportCmd)
	appsImportCmd.Flags().StringVar(&flogoWebHost, "host", "http://localhost:3303", "The URL for the Flogo Web UI")
	appsImportCmd.Flags().StringVar(&flogoFileName, "filename", "flogo.json", "The name of the file you want to import if you do not specify 'dir'")
	appsImportCmd.Flags().BoolVar(&readDir, "dir", false, "import all JSON files in the current directory")
}

// runAppsImport is the actual execution of the command
func runAppsImport(cmd *cobra.Command, args []string) {
	// Print flags
	fmt.Printf("\nThe URL for the Flogo Web UI URL has been set to: %s\n", flogoWebHost)
	if readDir {
		fmt.Printf("Importing all JSON files from                   : %s\n", util.CurrentDirectory())
	}
	fmt.Println("")

	// Prepare request URL
	flogoURL := fmt.Sprintf("%s%s", flogoWebHost, flogoAPIEndpoint)

	if readDir {
		// Read all the filenames in the current directory (does not traverse subdirectories)
		files, err := ioutil.ReadDir(".")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}

		// For each file that ends with .json try to import it
		foundFiles := false
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".json") {
				foundFiles = true
				fmt.Printf("Attempt to import: %s\n", file.Name())
				importToFlogoWeb(file.Name(), flogoURL)
			}
		}

		if !foundFiles {
			fmt.Printf("No JSON files have been found in                : %s\n\n", util.CurrentDirectory())
		}
	} else {
		// import a single file
		importToFlogoWeb(flogoFileName, flogoURL)
	}
}

// Function importToFlogoWeb() tries to import the file to the Flogo Web UI. The function
// takes two parameters:
// - fileName: the string which contains the name of the file
// - flogoURL: the URL of the Flogo Web UI, with the API endpoints
// If an error occurs, most likely because the Flogo Web UI doesn't have the activities
// needed, it will try to display which activities need to be imported
func importToFlogoWeb(fileName, flogoURL string) {
	// Read the file
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error while reading file %s\n", fileName)
		fmt.Println(err.Error())
	}

	// Prepare the API request
	req, err := http.NewRequest("POST", fmt.Sprint(flogoURL+flogoAppImport), bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	// Prepare the HTTP client
	client := &http.Client{}

	// Execute the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	// Print the response if not okay
	if resp.StatusCode != 200 {
		fmt.Printf("Couldn't import %s\n", fileName)
		fmt.Printf("  HTTP Status: %v\n", resp.StatusCode)

		var dat map[string]interface{}

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)

		byt := []byte(buf.String())
		if err := json.Unmarshal(byt, &dat); err != nil {
			fmt.Println(err.Error())
		}

		// Find the errors (most likely they're missing activities or triggers)
		strs := dat["errors"].([]interface{})[0].(map[string]interface{})["meta"].(map[string]interface{})["details"].([]interface{})
		for _, item := range strs {
			tempItem := item.(map[string]interface{})
			fmt.Printf("  %s\n", tempItem["message"])
		}
	} else {
		fmt.Printf("Succesfully imported %s\n", fileName)
	}
	fmt.Println("")
}
