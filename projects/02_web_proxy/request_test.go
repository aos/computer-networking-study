package main

import (
	"net"
	"net/url"
	"testing"
)

func TestCreateRequest(t *testing.T) {
	server, client := net.Pipe()
	defer client.Close()

	request := "GET www.google.com HTTP/1.0\r\n"

	want := Request{Method: "GET",
		ParsedURL: url.URL{
			Scheme: "http",
			Host:   "www.google.com:80",
		},
	}

	go func() {
		defer server.Close()
		got := *CreateRequest(server)

		if got.Method != want.Method {
			t.Errorf("Expected %q, got %q", want.Method, got.Method)
		}
		if got.ParsedURL.Scheme != want.ParsedURL.Scheme {
			t.Errorf("Expected %q, got %q",
				want.ParsedURL.Scheme, got.ParsedURL.Scheme,
			)
		}
		if got.ParsedURL.Host != want.ParsedURL.Host {
			t.Errorf("Expected %q, got %q",
				want.ParsedURL.Host, got.ParsedURL.Host,
			)
		}
	}()

	_, err := client.Write([]byte(request))
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}
