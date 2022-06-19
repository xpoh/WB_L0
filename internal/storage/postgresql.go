package storage

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/xpoh/WB_L0/internal/order"
	"os"
)

type PostgreSQl struct {
	dbUrl string
}

func NewPostgreSQl(db string) *PostgreSQl {
	return &PostgreSQl{dbUrl: db}
}

func (p *PostgreSQl) WriteAll(m map[string]order.Order) error {
	dbCfg := pgx.ConnConfig{
		Host: p.dbUrl,
	}
	conn, err := pgx.Connect(dbCfg)
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
	dbCfg := pgx.ConnConfig{
		Host: p.dbUrl,
	}
	conn, err := pgx.Connect(dbCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}
	defer conn.Close()

	data, _ := ord.SaveToJson()
	_, err = conn.Query("insert into wb_orders values ($1, $2)", id, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DB] add order failed: %v\n", err)
	}
	return err
}

func (p *PostgreSQl) Find(id string) (order.Order, error) {
	ord := order.Order{}

	dbCfg := pgx.ConnConfig{
		Host: p.dbUrl,
	}
	conn, err := pgx.Connect(dbCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	r := conn.QueryRow("select jsondata from wb_orders where order_uid=$1", id)
	err = r.Scan(&ord)
	if err != nil {
		return order.Order{}, err
	}
	return ord, nil
}

func (p *PostgreSQl) ReadAll() (map[string]order.Order, error) {
	m := make(map[string]order.Order)
	ord := order.Order{}
	var id string

	dbCfg := pgx.ConnConfig{
		Host: p.dbUrl,
	}
	conn, err := pgx.Connect(dbCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	rs, err := conn.Query("select id, jsondata from wb_orders")
	if err != nil {
		return m, err
	}

	for rs.Next() {
		rs.Scan(&id, &ord)
		m[id] = ord
	}

	return m, nil
}
