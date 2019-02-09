package main

import (
	"fmt"

	"github.com/r3labs/sse"
)

func main() {
	client := sse.NewClient("http://localhost:8080/events")
	client.Headers["X-Forwarded-For"] = "127.0.0.1"
	fmt.Println("connecting...")
	client.Subscribe("messages", func(msg *sse.Event) {
		fmt.Printf("\r%v", string(msg.Data))
	})
}
