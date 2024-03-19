package main

import (
	"fmt"
	"net/http"
	"os"
)

var apiURL string

func init() {
	// Retrieve API URL from environment variable
	apiURL = os.Getenv("API_URL")
	if apiURL == "" {
		fmt.Println("API_URL environment variable is not set.")
		os.Exit(1)
	} else {
		fmt.Printf("API URL: %s\n", apiURL)
	}
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
