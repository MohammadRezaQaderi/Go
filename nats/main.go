package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

const server = "nats://127.0.0.1:4222"

type Message struct {
	ID string
}

func main() {
	// Connect to NATS
	c, err := nats.Connect(server)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(c.ConnectedAddr())
	log.Print(c.DiscoveredServers())

	// Make end coding on Connection
	ec, err := nats.NewEncodedConn(c, nats.GOB_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	// make channel for sure that message is given by subscriber
	ch := make(chan struct{})

	if _, err := ec.Subscribe("hello-world", func(m *Message) {
		log.Print(m)
		close(ch)
	}); err != nil {
		log.Fatal(err)
	}

	// Publish Message on subject
	if err := ec.Publish("hello-world", Message{ID: "1378"}); err != nil {
		log.Fatal(err)
	}
	<-ch

}
