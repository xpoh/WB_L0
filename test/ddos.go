package main

import (
	"github.com/nats-io/stan.go"
	"github.com/xpoh/WB_L0/internal/order"
	"log"
)

func main() {
	order := order.NewOrder()

	clientID := "test_ddos_"
	URL := "0.0.0.0:4222"
	clusterID := "test-cluster"
	subj := "add_orders"

	sc, err := stan.Connect(clusterID, clientID,
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can'b connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, URL)
	}
	defer sc.Close()

	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", URL, clusterID, clientID)

	for i := 0; i < 10; i++ {
		if err = order.FillFakeData(); err == nil {
			if data, err := order.SaveToJson(); err == nil {
				err = sc.Publish(subj, data)
			}
		} else {
			log.Println(err)
		}
	}
}
