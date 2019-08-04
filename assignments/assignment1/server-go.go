package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments")
		fmt.Println("Usage: ./server-go $PORT")
		os.Exit(1)
	}

	port := os.Args[1]
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()

	}
}
