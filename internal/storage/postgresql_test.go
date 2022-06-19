package storage

import (
	"github.com/xpoh/WB_L0/internal/order"
	"reflect"
	"testing"
)

func TestPostgreSQl_Add(t *testing.T) {
	ord := order.Order{}
	ord.FillFakeData()

	type fields struct {
		dbUrl string
	}
	type args struct {
		id  string
		ord order.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Add one",
			fields: fields{dbUrl: "localhost:5432"},
			args: args{
				id:  ord.OrderUid,
				ord: ord,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostgreSQl{
				dbUrl: tt.fields.dbUrl,
			}
			if err := p.Add(tt.args.id, tt.args.ord); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostgreSQl_Find(t *testing.T) {
	type fields struct {
		dbUrl string
	}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostgreSQl{
				dbUrl: tt.fields.dbUrl,
			}
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
	type fields struct {
		dbUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]order.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostgreSQl{
				dbUrl: tt.fields.dbUrl,
			}
			got, err := p.ReadAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgreSQl_WriteAll(t *testing.T) {
	type fields struct {
		dbUrl string
	}
	type args struct {
		m map[string]order.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PostgreSQl{
				dbUrl: tt.fields.dbUrl,
			}
			if err := p.WriteAll(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("WriteAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
