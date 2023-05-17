package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Configuration holds the application version
type Configuration struct {
	Version string `json:"version"`
}

func main() {
	// Load the configuration file
	config, err := loadConfiguration("config.json")
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Set up the HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf("hello-app says hi! [version: %s]", config.Version)
		fmt.Fprintln(w, response)
	})

	// Start the server
	log.Println("Starting hello-app server...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

func loadConfiguration(filename string) (Configuration, error) {
	var config Configuration

	// Read the configuration file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	// Unmarshal the JSON data into the Configuration struct
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
