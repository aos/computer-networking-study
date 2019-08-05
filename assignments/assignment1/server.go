package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments\nUsage: ./server $PORT")
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
		go func(c net.Conn) {
			fmt.Printf("Client connected: %s\n", c.RemoteAddr())
			io.Copy(os.Stdout, c)
			c.Close()
		}(conn)
	}
}
