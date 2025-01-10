package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		return
	}

	username := os.Args[1]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status is not 200 OK
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch data. HTTP Status: %s\n", resp.Status)
		return
	}

	// Read and display the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	fmt.Println(string(body)) // Display raw JSON response
}
