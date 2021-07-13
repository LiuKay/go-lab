package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("start server...")

	ln, _ := net.Listen("tcp", ":8888")

	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		fmt.Fprintf(conn, message+"\n")
	}
}
