package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

// Request holds all parts of the request being sent
// optional
type Request struct {
	Method     string
	RequestURI string
	Headers    map[string]string
	Body       string
}

// CreateRequest creates a request struct with all the fields, given a Conn
//
// Request-line: GET / HTTP/1.0<CRLF>
// Headers: 0 or more
// <CRLF>
// Body
//
// So there will be no more than two CRLF. A body will only appear in a POST
// request. So after scanning the first line and seeing that it's not a POST,
// we can be sure to pass along the request if we see another CRLF.
func CreateRequest(c net.Conn) *Request {
	var req Request
	countCrlf := 0
	s := bufio.NewScanner(c)
	s.Split(ScanCRLF)

	for s.Scan() {
		countCrlf++
		// Get the request line
		if countCrlf == 1 {
			line := s.Text()
			received := strings.Split(line, " ")
			if len(received) < 3 {
				log.Fatalln("Incorrectly formatted request")
			}
			// TODO: Malformed request
			// TODO: Check for POST requests
			req.Method = received[0]
			req.RequestURI = received[1]
			// TODO: Parse headers
		}
		// Do other things (after second CRLF)
		// TODO: Parse body if POST
		break
	}
	return &req
}
