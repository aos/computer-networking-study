package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const recvBufferSize = 2048

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

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	buf := make([]byte, recvBufferSize)
	for {
		n, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Print(string(buf[:n]))
				break
			}
			log.Fatalln(err)
		}
		fmt.Print(string(buf[:n]))
	}
	c.Close()
}
