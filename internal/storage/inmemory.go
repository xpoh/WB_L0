package storage

import "github.com/xpoh/WB_L0/internal/order"

type Storager interface {
	Add(id string, ord order.Order) error
	Find(id string) (order.Order, error)
	ReadAll() (map[string]order.Order, error)
	WriteAll(map[string]order.Order, error)
}

type ErrorStoragerNotFind struct{}

func (e ErrorStoragerNotFind) Error() string {
	return "Not find error"
}

type InMemory struct {
	Orders map[string]order.Order
}

func (i *InMemory) ReadAll() (map[string]order.Order, error) {
	return i.Orders, nil
}

func (i *InMemory) WriteAll(m map[string]order.Order) error {
	i.Orders = m
	return nil
}

func (i *InMemory) Add(id string, ord order.Order) error {
	i.Orders[id] = ord
	return nil
}

func (i *InMemory) Find(id string) (order.Order, error) {
	if _, ok := i.Orders[id]; ok {
		return i.Orders[id], nil
	} else {
		return order.Order{}, ErrorStoragerNotFind{}
	}
}

func NewInMemory() *InMemory {
	m := &InMemory{}
	m.Orders = make(map[string]order.Order)
	return m
}
