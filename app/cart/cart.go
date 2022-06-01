package cart

import "github.com/jiramot/tdd/app/product"

type Cart struct {
	Items      []orderItem
	TotalPrice float64
}

type orderItem struct {
	Product  product.Product
	Quantity int
}

func NewCard() *Cart {
	return &Cart{TotalPrice: 0.0, Items: make([]orderItem, 0)}
}

func (c *Cart) Add(product product.Product, quantity int) *Cart {
	item := orderItem{Product: product, Quantity: quantity}
	c.Items = append(c.Items, item)
	c.TotalPrice += item.Product.Price * float64(item.Quantity)
	return c
}
