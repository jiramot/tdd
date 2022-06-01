package product

import "testing"

func TestGetProductPrice(t *testing.T) {
	ipad := NewProduct("ipad", 10000, 2)

	if ipad.Price != 10000 {
		t.Errorf("expect price %v but got %v", 10000, ipad.Price)
	}
}
