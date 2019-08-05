package main

import (
	"bufio"
	"fmt"
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
	s, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(s)
}
