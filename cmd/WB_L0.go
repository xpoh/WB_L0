package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"github.com/xpoh/WB_L0/internal/queue"
	"github.com/xpoh/WB_L0/internal/storage"
	"github.com/xpoh/WB_L0/internal/worker"
	"net/http"
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

	backup := storage.NewPostgreSQl("postgres://akaddr:aswqas@localhost:5432/base10")
	cache := storage.NewInMemory(backup)

	w := make([]*worker.Worker, queue.Cfg.MaxWorkers)
	for i, v := range w {
		v = worker.NewWorker("worker #"+fmt.Sprint(i), ctx, ch, cache)
		go v.Run()
	}

	// Create web server
	router := gin.Default()
	router.StaticFS("/website", http.Dir("./website"))

	router.GET("/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		if val, err := cache.Find(id); err == nil {
			c.JSON(http.StatusOK, val)
		} else {
			c.String(http.StatusNotFound, "Not find %s", id)
		}
	})

	go router.Run(":8080")

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Run cleanup when signal is received
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
	ctx.Done()
}
