package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	host := "127.0.0.1"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	protocol := "icmp"
	netaddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialIP("ip4:"+protocol, nil, netaddr)
	if err != nil {
		panic(err)
	}

	// Structure of an ICMP packet
	// []byte{
	//	0x08,		// ICMP type (echo request)
	//	0x00,		// ICMP subtype
	//	0x00, 0x00,	// Checksum (2 bytes)
	//	0x00, 0x00,	// Identifier (2 bytes)
	//	0x00, 0x00,	// Sequence number (2 bytes)
	//	0x00,		// (optional payload)
	//}
	// identifier := []byte{0x2C, 0xCC}
	go func(conn *net.IPConn) {
		buf := make([]byte, 1024)
		for {
			numRead, _, err := conn.ReadFrom(buf)
			if err != nil {
				panic(err)
			}
			fmt.Println("Number of bytes returned", numRead)
			fmt.Printf("% X\n", buf[:numRead])
		}
	}(conn)

	for {
		icmpPacket := []byte{
			0x08, 0x00,
			0x28, 0x3B,
			0x2C, 0xCC,
			0x00, 0x01,
			0x4F, 0x93, 0x65, 0x5E,
			0x00, 0x00, 0x00, 0x00,
			0x25, 0x33, 0x0A, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,
			0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F, 0x20, 0x21, 0x22, 0x23,
			0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D,
			0x2E, 0x2F, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
		}
		n, err := conn.Write(icmpPacket)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Written %d bytes to %s\n", n, conn.RemoteAddr().String())

		time.Sleep(1 * time.Second)
	}
}
