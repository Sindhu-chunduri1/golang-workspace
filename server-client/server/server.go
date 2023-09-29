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

	listenAddr := "localhost:8080"
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		// return
	}

	defer listener.Close()

	fmt.Println("Server is listening on", listenAddr)

	for {										
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	defer conn.Close()

	decoder := json.NewDecoder(conn)

	for {
		var req Request
		if err := decoder.Decode(&req); err != nil {	

			if err.Error() == "EOF" {
				fmt.Println("Client closed connection")
				break
			}

			fmt.Println("Error decoding JSON:", err)
			// return
		}

		// Process the request based on the operation
		var resp Response
		resp.Status = "OK"
		switch req.Operation {

		case "string":

			// Handle string data
			data, ok := req.Data.(string)
			if !ok {

				resp.Status = "Error"
				resp.Message = "Invalid data type for 'string' operation"
			} else {
				resp.Message = "Received string: " + data
			}

		case "file":

			// Handle file data (example)
			data, ok := req.Data.(string)
			if !ok {
				resp.Status = "Error"
				resp.Message = "Invalid data type for 'file' operation"
			} else {
				// Process the file (e.g., save it)
				// You can implement your file handling logic here
				resp.Message = "Received file: " + data
			}

		default:
			resp.Status = "Error"
			resp.Message = "Unknown operation"

		}

		encoder := json.NewEncoder(conn)
		if err := encoder.Encode(resp); err != nil {
			fmt.Println("Error encoding JSON response:", err)
			// return
		}

		// Send a newline character to delimit the response
		_, err := conn.Write([]byte("\n"))
		if err != nil {
			fmt.Println("Error sending delimiter:", err)
			// return
		}
	}
}
