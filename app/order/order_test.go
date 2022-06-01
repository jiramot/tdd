package order

import (
	"testing"

	"github.com/jiramot/tdd/app/cart"
	"github.com/jiramot/tdd/app/product"
	"github.com/jiramot/tdd/app/shipping"
)

func TestCheckOutOrder(t *testing.T) {
	ipad := product.NewProduct("ipad", 5000, 0.5)

	c := cart.NewCart()
	c.Add(ipad, 1)

	sh := shipping.NewShippingService(100)
	o := NewOrderService(sh)
	order := o.CheckOut(c)

	if order.OrderItem[0].Product.Name != "ipad" {
		t.Errorf("expect product %v but got %v", ipad.Name, order.OrderItem[0].Product.Name)
	}
}

func TestCheckOutShippingCost(t *testing.T) {
	ipad := product.NewProduct("ipad", 5000, 0.5)

	c := cart.NewCart()
	c.Add(ipad, 1)

	sh := shipping.NewShippingService(100)
	o := NewOrderService(sh)
	order := o.CheckOut(c)

	if order.ShippingCost != 50 {
		t.Errorf("expect product %v but got %v", 50, order.ShippingCost)
	}
}

func TestCheckOutGrandTotalPrice(t *testing.T) {
	ipad := product.NewProduct("ipad", 5000, 0.5)

	c := cart.NewCart()
	c.Add(ipad, 1)

	sh := shipping.NewShippingService(100)
	o := NewOrderService(sh)
	order := o.CheckOut(c)

	if order.GrandTotalPrice != 5050 {
		t.Errorf("expect product %v but got %v", 5050, order.GrandTotalPrice)
	}
}
