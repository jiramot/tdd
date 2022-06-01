package cart

import (
	"github.com/jiramot/tdd/app/product"
)

type Cart struct {
	Items       []orderItem
	RemovedItem []removeProductItem
}

type orderItem struct {
	Product  product.Product
	Quantity int
}

type removeProductItem struct {
	Product product.Product
}

func NewCart() *Cart {
	return &Cart{Items: make([]orderItem, 0), RemovedItem: make([]removeProductItem, 0)}
}

func (c *Cart) RemoveItem(product product.Product) *Cart {
	var newItems []orderItem
	for _, item := range c.Items {
		if item.Product.Name != product.Name {
			newItems = append(newItems, item)
		}
	}

	removeProductItems := removeProductItem{Product: product}
	c.RemovedItem = append(c.RemovedItem, removeProductItems)
	c.Items = newItems
	return c
}

func (c *Cart) Add(product product.Product, quantity int) *Cart {
	item := orderItem{Product: product, Quantity: quantity}
	c.Items = append(c.Items, item)
	return c
}
