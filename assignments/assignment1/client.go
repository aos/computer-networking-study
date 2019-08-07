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

	if len(os.Args) != 3 {
		fmt.Println("Wrong number of arguments\nUsage: ./client $IP_ADDR $PORT")
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", os.Args[1]+":"+os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}

	buf := make([]byte, recvBufferSize)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			if err == io.EOF {
				_, err = conn.Write(buf[:n])
				if err != nil {
					log.Fatalln(err)
				}
				break
			}
			log.Fatalln(err)
		}

		_, err = conn.Write(buf[:n])
		if err != nil {
			log.Fatalln(err)
		}
	}
	conn.Close()
}
