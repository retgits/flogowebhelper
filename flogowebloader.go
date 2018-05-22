// Package main is a utility to load Flogo apps into the Flogo Web UI using the commandline as
// opposed to using the import button. Especially when creating a new docker image, it helps to
// speed up the process of loading apps considerably.
//
// You can build an executable out of this app using the command
//  go build
// Usage of the app:
// -dir
//  Upload all JSON files in the current directory
// -filename string
//  The name of the file you want to upload if you do not specify 'dir' (default "flogo.json")
// -host string
//  The URL for Flogo Web UI (default "http://localhost:3303")
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// The API endpoint for Flogo Web UI
	flogoAPIEndpoint string = "/api/v2/"

	// The app import endpoint for Flogo Web UI
	flogoAppImport string = "apps:import"
)

// Function main() is main method of the app.
func main() {
	// Read the command line flags
	flogoHost := flag.String("host", "http://localhost:3303", "The URL for Flogo Web UI")
	readDir := flag.Bool("dir", false, "Upload all JSON files in the current directory")
	fileName := flag.String("filename", "flogo.json", "The name of the file you want to upload if you do not specify 'dir'")

	flag.Parse()

	// Prepare request URL
	flogoURL := fmt.Sprintf(*flogoHost + flogoAPIEndpoint)
	fmt.Printf("Setting Flogo Web UI URL to: %s\n", flogoURL)

	if *readDir {
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
		uploadToFlogoWeb(*fileName, flogoURL)
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
