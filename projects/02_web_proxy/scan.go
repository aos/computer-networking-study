package main

import (
	"bytes"
)

// ScanCRLF searches for a literal CRLF and returns that token
func ScanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\r', '\n'}); i >= 0 {
		// CRLF-terminated line
		return i + 2, data[0 : i+2], nil
	}
	// Final, non-terminated line
	if atEOF {
		return len(data), data, nil
	}
	// Request more data
	return 0, nil, nil
}
