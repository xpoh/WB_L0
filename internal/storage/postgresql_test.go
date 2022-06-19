package storage

import (
	"github.com/xpoh/WB_L0/internal/order"
	"log"
	"reflect"
	"testing"
)

func TestPostgreSQl_Add(t *testing.T) {
	type fields *PostgreSQl

	type args struct {
		id  string
		ord order.Order
	}

	p := NewPostgreSQl("postgres://akaddr:aswqas@localhost:5432/base10")

	ord := order.NewOrder()
	ord.FillFakeData()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Simple one add",
			fields: p,
			args: args{
				id:  ord.OrderUid,
				ord: *ord,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := p.Add(tt.args.id, tt.args.ord); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgreSQl_Find(t *testing.T) {
	type fields *PostgreSQl
	p := NewPostgreSQl("postgres://akaddr:aswqas@localhost:5432/base10")
	ord := order.NewOrder()
	ord.FillFakeData()

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    order.Order
		wantErr bool
	}{
		{
			name:    "Simple find test",
			fields:  p,
			args:    args{id: ord.OrderUid},
			want:    *ord,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p.Add(tt.args.id, *ord)
			got, err := p.Find(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgreSQl_ReadAll(t *testing.T) {
	type fields *PostgreSQl
	p := NewPostgreSQl("postgres://akaddr:aswqas@localhost:5432/base10")

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Simple Read all test",
			fields:  p,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := p.ReadAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println(got)
		})
	}
}
