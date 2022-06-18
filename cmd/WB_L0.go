package main

import (
	"context"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/xpoh/WB_L0/internal/queue"
	"github.com/xpoh/WB_L0/internal/worker"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	ch := make(chan stan.Msg, queue.Cfg.MaxWorkers)

	q := queue.NewQueue(ctx, ch)

	go func() {
		err := q.Start()
		if err != nil {
			fmt.Println(err)
		}
	}()

	w := make([]*worker.Worker, queue.Cfg.MaxWorkers)
	for i, v := range w {
		v = worker.NewWorker("worker #"+fmt.Sprint(i), ctx, ch)
		go v.Run()
	}

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Run cleanup when signal is received
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	for range signalChan {
		fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
		ctx.Done()
	}
}
