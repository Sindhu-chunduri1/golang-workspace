package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error in listening to port 8080: ", err)
	}
	defer listener.Close() //todo - what is defer ?
	fmt.Println("Server listening to port 8080")

	for { //todo - what kind of for loop is this (how many types of for loop in GoLang ?)
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in accepting connection :", err)
		}
		go handleClient(conn) // todo - read go routine
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024) //todo - read the data from connection - donot hard code
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error in reading ", err)
	}

	messagerecievedfromclient := string(buffer)
	fmt.Println("Message recieved from client :", messagerecievedfromclient)

	responsemessage := "Hello from server"
	_, err = conn.Write([]byte(responsemessage))
	if err != nil {
		fmt.Println("Error in writing ", err)
	}

}
