// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/retgits/flogowebhelper/util"
	"github.com/spf13/cobra"
)

// appsExportCmd represents the appsExport command
var appsExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports Flogo apps from the Flogo Web UI",
	Run:   runAppsExport,
}

// Flags
var ()

// Variables
var ()

// init registers the command and flags
func init() {
	appsCmd.AddCommand(appsExportCmd)
	appsExportCmd.Flags().StringVar(&flogoWebHost, "host", "http://localhost:3303", "The URL for the Flogo Web UI")
}

// runAppsExport is the actual execution of the command
func runAppsExport(cmd *cobra.Command, args []string) {
	// Print flags
	fmt.Printf("\nThe URL for the Flogo Web UI URL has been set to: %s\n\n", flogoWebHost)

	// Prepare request URL
	flogoURL := fmt.Sprintf("%s%s", flogoWebHost, flogoAPIEndpoint)

	// Get all the app IDs
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", flogoURL, flogoApps), nil)
	req.Header.Set("Content-Type", "application/json")

	// Prepare the HTTP client
	client := &http.Client{}

	// Execute the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Error while connecting to Flogo Web API\n")
		fmt.Printf("  HTTP Status: %v\n", resp.StatusCode)
		os.Exit(2)
	}

	// Create a directory for the exported app
	err = os.MkdirAll(filepath.Join(util.CurrentDirectory(), "apps"), os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	// Get the IDs from the apps
	var data map[string]interface{}
	if err := json.Unmarshal(readCloserToBytes(resp.Body), &data); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	for _, app := range data["data"].([]interface{}) {
		id := app.(map[string]interface{})["id"].(string)
		name := app.(map[string]interface{})["name"].(string)
		downloadApp(flogoURL, id, name)
	}

	if len(data) == 0 {
		fmt.Printf("Flogo Web doesn't have any apps to export...\n")
	}

	fmt.Println()
}

func downloadApp(URL string, id string, name string) {
	name = strings.ToLower(strings.Replace(name, " ", "_", -1))

	// Download the app
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/%s:export?appmodel=standard", URL, flogoApps, id), nil)
	req.Header.Set("Content-Type", "application/json")

	// Prepare the HTTP client
	client := &http.Client{}

	// Execute the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Cannot export app %s\n", name)
		fmt.Printf("  HTTP Status: %v\n", resp.StatusCode)
		return
	}

	err = ioutil.WriteFile(filepath.Join(util.CurrentDirectory(), "apps", fmt.Sprintf(name, ".json")), readCloserToBytes(resp.Body), os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Successfully exported %s\n", name)
}

func readCloserToString(rc io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	return buf.String()
}

func readCloserToBytes(rc io.ReadCloser) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	return buf.Bytes()
}
