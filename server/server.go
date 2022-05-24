package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	fmt.Println("Starting server...")

	ln, _ := net.Listen("tcp", "localhost:8080")

	fmt.Println("Server started! Waiting for connection")

	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received: ", string(message))
	}
}
