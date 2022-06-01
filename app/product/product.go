package product

type Product struct {
	ID     string
	Name   string
	Price  float64
	Weight float64
}

func NewProduct(name string, price float64, weight float64) Product {
	return Product{Name: name, Price: price, Weight: weight}
}
