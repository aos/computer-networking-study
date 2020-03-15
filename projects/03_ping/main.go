package main

import (
	"encoding/binary"
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
	defer conn.Close()

	// Structure of an ICMP packet
	// []byte{
	//	0x08, 0x00,	// ICMP type (echo request), subtype
	//	0x00, 0x00,	// Checksum (2 bytes)
	//	0x00, 0x00,	// Identifier (2 bytes)
	//	0x00, 0x00,	// Sequence number (2 bytes)
	//	0x00, ...	// (optional payload)
	//}

	identifier := uint16(0x0A41)
	go ReadEchoReply(conn, identifier)

	seqNum := uint16(1)
	fmt.Printf("PING %s (%s) %d(%d) bytes of data.\n",
		host, conn.RemoteAddr().String(), 16, 32,
	)
	for {
		icmpPacket := AssembleICMP(
			8, 0, identifier, seqNum, time.Now().UnixNano(),
		)

		_, err := conn.Write(icmpPacket)
		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)
		seqNum++
	}
}

// ReadEchoReply reads the echo reply from the sent packet
func ReadEchoReply(conn *net.IPConn, identifier uint16) {
	buf := make([]byte, 1024)
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		numRead, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}

		// Extract ICMP reply from IP packet
		icmpReply := buf[20:numRead]

		if identifier != binary.BigEndian.Uint16(icmpReply[4:6]) {
			fmt.Println("Incorrect identifer!")
		}

		timePacket := int64(binary.BigEndian.Uint64(icmpReply[8:16]))
		timeNano := time.Since(time.Unix(0, timePacket)).Nanoseconds()
		fmt.Printf(
			"-- %d bytes from %s: icmp_seq=%d time=%.2f ms\n",
			len(icmpReply),
			conn.RemoteAddr().String(),
			binary.BigEndian.Uint16(icmpReply[6:8]),
			float64(timeNano)/1000000,
		)
	}
}
