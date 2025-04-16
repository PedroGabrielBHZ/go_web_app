// Package main is the entry point of the application.
package main

import (
	"fmt"
	"net/http"
)

// HomeHandler handles requests to the root URL and responds with a welcome message.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

// StartServer initializes the HTTP server and starts listening on port 8080.
func StartServer() {
	http.HandleFunc("/", HomeHandler)

	// Print debug information about the server.
	fmt.Println("Server started on :8080")
	fmt.Println("Press Ctrl+C to stop the server")

	// Start the HTTP server.
	http.ListenAndServe(":8080", nil)
}

// main is the entry point of the application.
func main() {
	StartServer()
}
