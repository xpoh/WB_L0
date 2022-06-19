package storage

import (
	"github.com/xpoh/WB_L0/internal/order"
	"log"
	"sync"
)

type Storager interface {
	Add(id string, ord order.Order) error
	Find(id string) (order.Order, error)
	ReadAll() (map[string]order.Order, error)
	WriteAll(map[string]order.Order) error
}

type ErrorStoragerNotFind struct{}

func (e ErrorStoragerNotFind) Error() string {
	return "Not find error"
}

type InMemory struct {
	backup Storager
	Orders map[string]order.Order
	mux    sync.Mutex
}

func (i *InMemory) ReadAll() (map[string]order.Order, error) {
	return i.Orders, nil
}

func (i *InMemory) WriteAll(m map[string]order.Order) error {
	i.mux.Lock()
	defer i.mux.Unlock()

	i.Orders = m
	err := i.backup.WriteAll(m)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (i *InMemory) Add(id string, ord order.Order) error {
	i.mux.Lock()
	defer i.mux.Unlock()

	i.Orders[id] = ord
	err := i.backup.Add(id, ord)
	return err
}

func (i *InMemory) Find(id string) (order.Order, error) {
	if _, ok := i.Orders[id]; ok {
		return i.Orders[id], nil
	} else {
		return order.Order{}, ErrorStoragerNotFind{}
	}
}

func NewInMemory(backup Storager) *InMemory {
	m := &InMemory{
		backup: nil,
		Orders: nil,
	}
	m.Orders = make(map[string]order.Order)
	all, err := backup.ReadAll()
	if err != nil {
		log.Println(err)
		return nil
	}
	m.Orders = all
	m.backup = backup
	return m
}
