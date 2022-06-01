package order

import (
	"github.com/jiramot/tdd/app/cart"
	"github.com/jiramot/tdd/app/product"
	"github.com/jiramot/tdd/app/shipping"
)

type Order struct {
	OrderItem       []orderItem
	TotalPrice      float64
	ShippingCost    float64
	GrandTotalPrice float64
}

type OrderService struct {
	Scs *shipping.ShippingService
}

func NewOrderService(scs *shipping.ShippingService) *OrderService {
	return &OrderService{Scs: scs}
}

type orderItem struct {
	Product  product.Product
	Quantity int
}

func (os *OrderService) CheckOut(cart *cart.Cart) *Order {
	o := &Order{OrderItem: make([]orderItem, 0)}
	for _, cartItem := range cart.Items {
		item := orderItem{Product: cartItem.Product, Quantity: cartItem.Quantity}
		o.OrderItem = append(o.OrderItem, item)
		o.TotalPrice += cartItem.Product.Price * float64(cartItem.Quantity)
		o.ShippingCost += os.Scs.CalculateShippingCost(cartItem.Product.Weight) * float64(cartItem.Quantity)
	}
	o.GrandTotalPrice += o.TotalPrice + o.ShippingCost
	return o
}
