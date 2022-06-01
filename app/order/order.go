package order

import "github.com/jiramot/tdd/app/product"

type Order struct {
	OrderItem  []orderItem
	TotalPrice float64
}

type orderItem struct {
	Product  product.Product
	Quantity int
}

func NewOrder() *Order {
	return &Order{OrderItem: make([]orderItem, 0), TotalPrice: 0}
}

func (o *Order) Add(product product.Product, quantity int) *Order {
	item := orderItem{Product: product, Quantity: quantity}
	o.OrderItem = append(o.OrderItem, item)
	o.TotalPrice += product.Price * float64(quantity)
	return o
}
