package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	// Connect to NATS
	nc, _ := nats.Connect(nats.DefaultURL)

	// Create JetStream Context
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})
	
	// Simple Stream Publisher
	js.Publish("ORDERS.scratch", []byte("hello"))

	// Simple Async Stream Publisher
	for i := 0; i < 500; i++ {
		js.PublishAsync("ORDERS.scratch", []byte("hello"))
	}
	select {
	case <-js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}

	// Simple Async Ephemeral Consumer
	js.Subscribe("ORDERS.*", func(m *nats.Msg) {
		fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
	})

	// Simple Pull Consumer
	sub, err := js.PullSubscribe("ORDERS.*", "MONITOR")
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := sub.Fetch(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("msgs: %+v", msgs)
	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()
}
