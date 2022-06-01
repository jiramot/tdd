package cart

import (
	"testing"

	"github.com/jiramot/tdd/app/product"
)

func TestSumTotalPriceWhenAddNewProductToCart(t *testing.T) {
	ipad := product.NewProduct("ipad", 10000)
	orderItem := OrderItem{Product: ipad, Quantity: 1}

	c := NewCard()
	c.Add(orderItem)

	if c.TotalPrice != 10000 {
		t.Errorf("expect total price %v but got %v", 10000, c.TotalPrice)
	}

}

func TestAddNewProductToCart(t *testing.T) {
	ipad := product.NewProduct("ipad", 10000)
	orderItem := OrderItem{Product: ipad, Quantity: 1}

	c := NewCard()
	c.Add(orderItem)

	if c.Items[0].Product.Name != "ipad" {
		t.Errorf("expect first product name %v but got %v", "ipad", c.Items[0].Product.Name)
	}

}

func TestAddAppleWatchToCart(t *testing.T) {
	appleWatch := product.NewProduct("apple watch", 10000)
	orderItem := OrderItem{Product: appleWatch, Quantity: 1}

	c := NewCard()
	c.Add(orderItem)
	if len(c.Items) != 1 {
		t.Errorf("expect no product in cart %v but got %v", 1, len(c.Items))
	}
}

func TestAddRiceCookerTwoQuantityToCart(t *testing.T) {
	riceCooker := product.NewProduct("rice cooker", 5000)
	orderItem := OrderItem{Product: riceCooker, Quantity: 2}

	c := NewCard()
	c.Add(orderItem)
	if len(c.Items) != 1 {
		t.Errorf("expect no product in cart %v but got %v", 1, len(c.Items))
	}
	if c.Items[0].Quantity != 2 {
		t.Errorf("expect quantity %v but got %v", 2, c.Items[0].Quantity)
	}
}
