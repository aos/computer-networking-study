package main

import (
	"bufio"
	"log"
	"net"
	"net/url"
	"strings"
)

// Request holds all parts of the request being sent
// The parsed URL is of type URL, which holds scheme, host, path, query
type Request struct {
	Method    string
	ParsedURL url.URL
	Headers   map[string]string
	Body      string
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

			// We need to add a scheme to the URL if it doesn't
			// have one because url.Parse() will not parse the
			// hostname correctly (as per RFC 3986)
			if !strings.HasPrefix(received[1], "//") {
				received[1] = "http://" + received[1]
			}

			parsedURL, err := url.Parse(received[1])
			if err != nil {
				log.Fatalln("Incorrectly formatted URL")
			}
			req.ParsedURL = *parsedURL

			// Default to absolute path if none given
			if req.ParsedURL.Path == "" {
				req.ParsedURL.Path = "/"
			}
			// Rebuild host to include port 80 if missing
			if req.ParsedURL.Port() == "" {
				req.ParsedURL.Host = req.ParsedURL.Host + ":80"
			}
		}
		// TODO: Parse headers
		// Do other things (after second CRLF)
		// TODO: Parse body if POST
		break
	}
	return &req
}
