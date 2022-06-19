package storage

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/xpoh/WB_L0/internal/order"
	"os"
)

type PostgreSQl struct {
	uri string
	cfg pgx.ConnConfig
}

func NewPostgreSQl(uri string) *PostgreSQl {
	cfg, err := pgx.ParseURI(uri)

	if err != nil {
		return nil
	}

	return &PostgreSQl{
		uri: uri,
		cfg: cfg,
	}
}

func (p *PostgreSQl) WriteAll(m map[string]order.Order) error {
	conn, err := pgx.Connect(p.cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	for id, ord := range m {
		_, err := conn.Query("insert into wb_orders values ($1, $2)", id, ord)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PostgreSQl) Add(id string, ord order.Order) error {
	conn, err := pgx.Connect(p.cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}
	defer conn.Close()

	data, _ := ord.SaveToJson()
	_, err = conn.Query("insert into wb_l0.wb_orders values ($1, $2)", id, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DB] add order failed: %v\n", err)
	}
	return err
}

func (p *PostgreSQl) Find(id string) (order.Order, error) {
	ord := order.Order{}
	var b []byte

	conn, err := pgx.Connect(p.cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	r := conn.QueryRow("select jsondata from wb_l0.wb_orders where order_uid=$1", id)
	err = r.Scan(&b)
	if err != nil {
		return order.Order{}, err
	}
	err = ord.LoadFromJson(b)
	if err != nil {
		return order.Order{}, err
	}
	if err != nil {
		return order.Order{}, err
	}
	return ord, nil
}

func (p *PostgreSQl) ReadAll() (map[string]order.Order, error) {
	m := make(map[string]order.Order)
	ord := order.Order{}
	var id string
	var b []byte

	conn, err := pgx.Connect(p.cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	rs, err := conn.Query("select order_uid, jsondata from wb_l0.wb_orders")
	if err != nil {
		return m, err
	}

	for rs.Next() {
		err := rs.Scan(&id, &b)
		if err != nil {
			return nil, err
		}
		err = ord.LoadFromJson(b)
		if err != nil {
			return nil, err
		}
		m[id] = ord
	}

	return m, nil
}
