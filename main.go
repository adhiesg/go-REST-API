package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func searchAnimeByNameHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the param is exist
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	name := r.Form.Get("name")
	url := apiURL + "/anime?sfw=true&q=" + name
	fmt.Println(url)

	// API call
	response, errCall := http.Get(url)
	if errCall != nil {
		http.Error(w, fmt.Sprintf("Error calling API: %v", errCall), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Check the status code of the response
	if response.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("API returned status code %d", response.StatusCode), http.StatusInternalServerError)
		return
	}

	// Read response body
	body, errBody := io.ReadAll(response.Body)
	if errBody != nil {
		http.Error(w, fmt.Sprintf("Error reading response body: %v", errBody), http.StatusInternalServerError)
		return
	}

	// Decode JSON response
	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding response: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/search", searchAnimeByNameHandler)

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
