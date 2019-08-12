// A web proxy built using RFC1945
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Proxy starts and listens to connections on specified $PORT
// The proxy listens for properly formatted HTTP requests from the client
// Each client is handled separately
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

// 1. Create the client's request
// 2. Parse the URL into: host, port, requested path
// 3. Connect to remote server using host:port
// 4. Send a properly formatted HTTP request using requested path
// 5. Return results to client
func handleConnection(c net.Conn) {
	defer c.Close()

	fmt.Printf("Connected to client: %s\n", c.RemoteAddr())
	req := *CreateRequest(c)

	reqConn, err := net.Dial("tcp", req.ParsedURL.Host)
	defer reqConn.Close()

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(reqConn, fmt.Sprintf("%s %v HTTP/1.0\r\n\r\n", req.Method, req.ParsedURL.Path))
	_, err = io.Copy(os.Stdout, reqConn)
	if err != nil {
		log.Fatalln(err)
	}
}
