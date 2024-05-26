package main

type Wine struct {
	Products
	Kind string
	Year int
}

func (w Wine) calculatePrice() int64 {
	var alcoholTax float64

	alcoholTax = float64(w.Price) * 0.23
	stateLiquorTax := float64(w.Price) * 0.05
	return w.Price + int64(alcoholTax) + int64(stateLiquorTax)
}
