// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// webloaderCmd represents the xxx command
var webloaderCmd = &cobra.Command{
	Use:   "webloader",
	Short: "Loads Flogo apps into the Flogo Web UI",
	Run:   runWebloader,
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
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(webloaderCmd)
	webloaderCmd.Flags().StringVar(&flogoWebHost, "host", "http://localhost:3303", "The URL for Flogo Web UI")
	webloaderCmd.Flags().StringVar(&flogoFileName, "filename", "flogo.json", "The name of the file you want to upload if you do not specify 'dir'")
	webloaderCmd.Flags().BoolVar(&readDir, "dir", false, "Upload all JSON files in the current directory")
}

// runWebloader is the actual execution of the command
func runWebloader(cmd *cobra.Command, args []string) {
	// Prepare request URL
	flogoURL := fmt.Sprintf("%s%s", flogoWebHost, flogoAPIEndpoint)
	fmt.Printf("Setting Flogo Web UI URL to: %s\n", flogoURL)

	if readDir {
		// Read all the filenames in the current directory (does not traverse subdirectories)
		files, err := ioutil.ReadDir(".")
		if err != nil {
			fmt.Println(err.Error())
		}

		// For each file that ends with .json try to upload it
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".json") {
				fmt.Printf("Attempt to upload: %s\n", file.Name())
				uploadToFlogoWeb(file.Name(), flogoURL)
			}
		}
	} else {
		// Upload a single file
		uploadToFlogoWeb(flogoFileName, flogoURL)
	}
}

// Function uploadToFlogoWeb() tries to upload the file to the Flogo Web UI. The function
// takes two parameters:
// - fileName: the string which contains the name of the file
// - flogoURL: the URL of the Flogo Web UI, with the API endpoints
// If an error occurs, most likely because the Flogo Web UI doesn't have the activities
// needed, it will try to display which activities need to be uploaded
func uploadToFlogoWeb(fileName, flogoURL string) {
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
		fmt.Printf("Couldn't upload %s\n", fileName)
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
		fmt.Printf("Succesfully uploaded %s\n", fileName)
	}
	fmt.Println("")
}
