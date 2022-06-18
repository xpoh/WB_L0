package worker

import (
	"context"
	"github.com/nats-io/stan.go"
	"github.com/xpoh/WB_L0/internal/order"
	"log"
)

type Worker struct {
	uid   string
	ctx   context.Context
	inQue chan stan.Msg
}

func NewWorker(uid string, ctx context.Context, in_que chan stan.Msg) *Worker {
	return &Worker{ctx: ctx, inQue: in_que, uid: uid}
}

func (w *Worker) Run() {
	var msg stan.Msg
	order := order.NewOrder()

	for {
		select {
		case <-w.ctx.Done():
			break
		case msg = <-w.inQue:
			if err := order.LoadFromJson(msg.Data); err != nil {
				log.Println(err)
			} else {
				log.Printf("Parsed order succseful [%s]. OrderId=%v\n", w.uid, order.OrderUid)
			}
		}
	}
}
