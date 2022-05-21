package main

import (
	"net"
	
)



func main() {
	
	ln, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		return
	}

	for {
	conn, err := ln.Accept()
	if err != nil{
		return
	}
	}
}
