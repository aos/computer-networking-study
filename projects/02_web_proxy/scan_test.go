package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestScanCRLF(t *testing.T) {
	simpleRequest := "GET / HTTP/1.0\r\n"
	buf := strings.NewReader(simpleRequest)
	s := bufio.NewScanner(buf)
	s.Split(ScanCRLF)

	for s.Scan() {
		want := s.Text()
		if want != simpleRequest {
			t.Errorf("Expected %q, got %q", simpleRequest, want)
		}
	}
}
