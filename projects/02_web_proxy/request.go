package main

import (
	"bufio"
	"log"
	"net"
	"net/url"
	"strconv"
	"strings"
)

var supportedMethods = map[string]bool{
	"GET":  true,
	"HEAD": true,
	"POST": true,
}

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
		receivedText := s.Text()
		// Parse the Request-line
		if countCrlf == 1 {
			received := strings.Split(receivedText, " ")
			if len(received) < 3 {
				log.Fatalln("Incorrectly formatted request")
			}
			req.Method = received[0]
			if _, ok := supportedMethods[req.Method]; !ok {
				log.Fatalln("Unsupported method:", req.Method)
			}

			// We need to add a scheme to the URL if it doesn't
			// have one because url.Parse() will not parse the
			// hostname correctly (as per RFC 3986, section 3)
			if !strings.HasPrefix(received[1], "//") {
				received[1] = "https://" + received[1]
			}

			parsedURL, err := url.Parse(received[1])
			if err != nil {
				log.Fatalln("Incorrectly formatted URL", err)
			}
			req.ParsedURL = *parsedURL

			// Default to absolute path if none given
			// RFC 1945 (section 5.1.2)
			if req.ParsedURL.Path == "" {
				req.ParsedURL.Path = "/"
			}
			// Rebuild host to include port 80 if missing
			if req.ParsedURL.Port() == "" {
				req.ParsedURL.Host = req.ParsedURL.Host + ":80"
			}
			continue
		}
		if countCrlf == 2 && receivedText != "\r\n" {
			req.Headers = createHeaderMap(receivedText[:len(receivedText)-2])
			continue
		}
		// The Content-Length header is required 100% in POST requests
		if req.Method == "POST" && receivedText != "\r\n" {
			bodyLength := len(receivedText) - 2
			req.Body = receivedText[:bodyLength]
			req.Headers["Content-Length"] = strconv.Itoa(bodyLength)
		}
		break
	}
	return &req
}

func createHeaderMap(headerString string) map[string]string {
	splitHeader := strings.Split(headerString, "\n")
	headerMap := make(map[string]string)
	for _, header := range splitHeader {
		if strings.Contains(header, ":") {
			key, value := splitColonHeader(header)
			headerMap[key] = value
		}
	}
	return headerMap
}

func splitColonHeader(header string) (string, string) {
	split := strings.SplitN(header, ":", 2)
	return split[0], split[1]
}
