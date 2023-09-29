package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error in connecting ", err)
	}
	defer conn.Close()

	message := "hello"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error in writing ", err)
	}

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error in reading ", err)
	}

	messagerecievedfromserver := string(buffer)
	fmt.Println("Message recieved from server : ", messagerecievedfromserver)
}
