package queue

import (
	"context"
	"github.com/nats-io/stan.go"
	"log"
)

type Queue struct {
	ctx   context.Context
	chOut chan stan.Msg
}

func NewQueue(ctx context.Context, chOut chan stan.Msg) *Queue {
	return &Queue{ctx: ctx, chOut: chOut}
}

type config struct {
	URL        string
	clusterID  string
	clientID   string
	subj       string
	qgroup     string
	durable    string
	MaxWorkers int
}

var Cfg = config{
	URL:        "0.0.0.0:4222",
	clusterID:  "test-cluster",
	clientID:   "test_client",
	subj:       "add_orders",
	qgroup:     "WB_L0",
	durable:    "durable",
	MaxWorkers: 100,
}

func printMsg(m *stan.Msg, i int) {
	log.Printf("[#%d] Received: %s\n", i, m)
}

func (q *Queue) Start() error {

	sc, err := stan.Connect(Cfg.clusterID, Cfg.clientID,
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, Cfg.URL)
	}
	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", Cfg.URL, Cfg.clusterID, Cfg.clientID)

	// Process Subscriber Options.
	startOpt := stan.StartWithLastReceived()

	mcb := func(msg *stan.Msg) {
		q.chOut <- *msg
	}

	sub, err := sc.QueueSubscribe(Cfg.subj, Cfg.qgroup, mcb, startOpt, stan.DurableName(Cfg.durable))
	if err != nil {
		err := sc.Close()
		if err != nil {
			return err
		}
		log.Fatal(err)
	}

	log.Printf("Listening on [%s], clientID=[%s], qgroup=[%s] durable=[%s]\n", Cfg.subj, Cfg.clientID, Cfg.qgroup, Cfg.durable)

	<-q.ctx.Done()

	err = sub.Unsubscribe()
	if err != nil {
		return err
	}
	err = sc.Close()
	if err != nil {
		return err
	}

	return err
}
