package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	serverURL := "http://localhost:8080"
	// Sending a GET request to the server
	resp, err := http.Get(serverURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response from server:", string(body))
}
