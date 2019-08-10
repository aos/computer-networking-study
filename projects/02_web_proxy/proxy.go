// A web proxy built using RFC1945
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments\nUsage: ./proxy $PORT")
		os.Exit(1)
	}

	ln, err := net.Listen("tcp", ":"+os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	fmt.Printf("Connected to client: %s\n", c.RemoteAddr())
	req := CreateRequest(c)

	fmt.Printf("Request received:\n%v\n", req)
}
