package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error in listening to port ", err)
	}
	defer listener.Close()

	fmt.Println("Server listening to port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in accepting connection ", err)
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	var receivedMessage Message
	err := decoder.Decode(&receivedMessage)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	//fmt.Println("Received Message : ", receivedMessage.Text)

	// Format the JSON response for better readability
	formattedResponse, err := json.MarshalIndent(receivedMessage, "", "    ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
	}

	fmt.Println("Received JSON response:")
	fmt.Println(string(formattedResponse))

	response := Message{Text: "Hello from server"}
	err = encoder.Encode(response)
	if err != nil {
		fmt.Println("Error in encoding JSON ", err)
	}
}
