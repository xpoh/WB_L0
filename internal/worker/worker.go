package worker

import (
	"context"
	"github.com/nats-io/stan.go"
	"github.com/xpoh/WB_L0/internal/order"
	"github.com/xpoh/WB_L0/internal/storage"
	"log"
)

type Worker struct {
	uid     string
	ctx     context.Context
	inQue   chan stan.Msg
	storage storage.Storager
}

func NewWorker(uid string, ctx context.Context, inQue chan stan.Msg, storager storage.Storager) *Worker {
	return &Worker{ctx: ctx, inQue: inQue, uid: uid, storage: storager}
}

func (w *Worker) Run() {
	var msg stan.Msg
	ord := order.NewOrder()

	for {
		select {
		case <-w.ctx.Done():
			return
		case msg = <-w.inQue:
			if err := ord.LoadFromJson(msg.Data); err != nil {
				log.Println(err)
			}
			err := w.storage.Add(ord.OrderUid, *ord)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
