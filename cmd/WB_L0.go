package main

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
)

func main() {
	clusterID := "test-cluster"
	clientID := "iddd"
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}
	// Simple Synchronous Publisher
	sc.Publish("foo", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming

	// Simple Async Subscriber
	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Unsubscribe
	sub.Unsubscribe()

	// Close connection
	sc.Close()

	sc, _ = stan.Connect("test-cluster", "clientid")

	// Create a queue subscriber on "foo" for group "bar"
	qsub1, _ := sc.QueueSubscribe("foo", "bar", qcb)

	// Add a second member
	qsub2, _ := sc.QueueSubscribe("foo", "bar", qcb)

	// Notice that you can have a regular subscriber on that subject
	sub, _ := sc.Subscribe("foo", cb)

	// A message on "foo" will be received by sub and qsub1 or qsub2.
}
