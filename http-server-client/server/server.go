package main

import (
	"fmt"
	"net/http"
)

func main() {

	//This function handling the incoming HTTP requests tothe path("/") where it responds with "Hello client"
	// This is entry point for program
	http.HandleFunc("/", handleRequest)
	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(port, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, client!")
}
