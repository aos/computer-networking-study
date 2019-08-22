package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

var addr = "192.168.1.8"

func main() {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/stats"}
	log.Printf("connecting to %s", u.String())

	c, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err, resp.StatusCode)
	}
	defer c.Close()

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
}
