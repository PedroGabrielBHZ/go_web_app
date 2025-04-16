package main

import (
	"app/pkg/handlers"
	"fmt"
	"net/http"
)

const port = ":8080"

func StartServer() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	fmt.Println("Starting server on port", port)
	fmt.Println("Access the server at: http://localhost" + port)
	fmt.Println("Press Ctrl+C to stop the server")

	http.ListenAndServe(port, nil)
}

func main() {
	StartServer()
}
