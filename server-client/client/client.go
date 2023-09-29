package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Request struct {
	Operation string      `json:"operation"`
	Data      interface{} `json:"data"`
} 

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {

	serverAddr := "localhost:8080"
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	defer conn.Close()

	// Send a string message
	reqString := Request{
		Operation: "string",
		Data:      "Hello !",
	}

	sendMessage(conn, reqString)

	// Send a file message (example)
	reqFile := Request{
		Operation: "file",
		Data:      "server-client/example.txt",
	}

	sendMessage(conn, reqFile)
}

func sendMessage(conn net.Conn, req Request) {

	//This code encodes the req request as JSON and sends it to the server using the encoder.
	//If there is an error during encoding for example if the request cannot be converted to JSON),
	//it prints an error message and returns from the function.

	encoder := json.NewEncoder(conn) // gave encoder to encode the req through network connection.
	if err := encoder.Encode(req); err != nil {
		fmt.Println("Error encoding JSON request:", err)
		return
	}

	// Read the response until a newline character to handle delimiters
	buffer := make([]byte, 1024)

	var response string

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error in reading :", err)
			// return
		}

		response += string(buffer[:n])
		if response[len(response)-1] == '\n' {
			break // Response is complete
		}
	}

	// Remove the trailing newline character
	response = response[:len(response)-1]

	// Decode and print the server response in JSON format
	var resp Response

	if err := json.Unmarshal([]byte(response), &resp); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		// return
	}

	responseJSON, _ := json.Marshal(resp)
	fmt.Println("Server response:", string(responseJSON))
}
