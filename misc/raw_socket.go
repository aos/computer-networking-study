package main

import (
	"fmt"
	"net"
)

func main() {
	protocol := "tcp"
	netaddr, _ := net.ResolveIPAddr("ipv4", "127.0.0.1")
	conn, _ := net.ListenIP("ip4:"+protocol, netaddr)

	buf := make([]byte, 1024)
	numRead, _, _ := conn.ReadFrom(buf)
	fmt.Printf("% X\n", buf[:numRead])
}
