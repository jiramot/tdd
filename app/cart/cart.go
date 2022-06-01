package cart

import "github.com/jiramot/tdd/app/product"

type Cart struct {
	Items      []OrderItem
	TotalPrice float64
}

type OrderItem struct {
	Product  product.Product
	Quantity int
}

func NewCard() *Cart {
	return &Cart{TotalPrice: 0.0, Items: make([]OrderItem, 0)}
}

func (c *Cart) Add(item OrderItem) *Cart {
	c.Items = append(c.Items, item)
	c.TotalPrice += item.Product.Price * float64(item.Quantity)
	return c
}
