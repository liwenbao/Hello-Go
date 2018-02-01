package demo

import (
	"git.kkcoding.com/kk-golang/kk-lib/db"
)

type Product struct {
	db.Object
	Name  string
	Price float32
}

func (o *Product) GetName() string {
	return "Product"
}

func (o *Product) GetTitle() string {
	return "产品"
}

type Transaction struct {
	db.Object
	ProductId int64
	Count     int
	Totle     float32
}

func (o *Transaction) GetName() string {
	return "Transaction"
}

func (o *Transaction) GetTitle() string {
	return "交易"
}

func NewTransaction(p *Product, c int) *Transaction {
	t := &Transaction{}
	t.ProductId = p.GetId()
	t.Count = c
	t.Totle = float32(c) * p.Price
	return t
}
