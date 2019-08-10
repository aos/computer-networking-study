package main

import (
	"net"
	"testing"
)

func TestCreateRequest(t *testing.T) {
	server, client := net.Pipe()
	defer client.Close()

	request := "GET / HTTP/1.0\r\n"
	want := Request{Method: "GET", RequestURI: "/"}

	go func() {
		defer server.Close()
		got := *CreateRequest(server)

		if got.Method != want.Method {
			t.Errorf("Expected %q, got %q", want.Method, got.Method)
		}
		if got.RequestURI != want.RequestURI {
			t.Errorf("Expected %q, got %q", want.RequestURI, got.RequestURI)
		}
	}()

	_, err := client.Write([]byte(request))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}
