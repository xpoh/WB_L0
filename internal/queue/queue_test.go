package queue

import (
	"context"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/xpoh/WB_L0/internal/order"
	"github.com/xpoh/WB_L0/internal/worker"
	"log"
	"testing"
	"time"
)

var q *Queue
var ctx context.Context

func BenchmarkQueue_Start(b *testing.B) {
	order := order.NewOrder()

	if q == nil {
		ctx = context.Background()
		ch := make(chan stan.Msg, Cfg.MaxWorkers)

		q = NewQueue(ctx, ch)
		go func() {
			err := q.Start()
			if err != nil {
				return
			}
		}()

		w := make([]*worker.Worker, Cfg.MaxWorkers)
		for i, v := range w {
			v = worker.NewWorker("worker #"+fmt.Sprint(i), ctx, ch)
			go v.Run()
		}

		time.Sleep(time.Second)
	}

	Cfg.clientID = "test_id_" + fmt.Sprint(b.N)

	sc, err := stan.Connect(Cfg.clusterID, Cfg.clientID,
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can'b connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, Cfg.URL)
	}
	defer sc.Close()

	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", Cfg.URL, Cfg.clusterID, Cfg.clientID)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		if err = order.FillFakeData(); err == nil {
			if data, err := order.SaveToJson(); err == nil {
				err = sc.Publish(Cfg.subj, data)
			}
		} else {
			log.Println(err)
		}
		b.StopTimer()
	}
}
