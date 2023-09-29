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
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error in connecting to port ", err)
	}
	defer conn.Close()

	message := Message{Text: "Hello from client"}
	encoder := json.NewEncoder(conn)
	err = encoder.Encode(message)
	if err != nil {
		fmt.Println("Error in encoding JSON:", err)
		return
	}

	decoder := json.NewDecoder(conn)
	var response Message
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	//fmt.Println("Received message ", response.Text)

	// Format the JSON response for better readability
	formattedResponse, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println("Received JSON response:")
	fmt.Println(string(formattedResponse))
}
