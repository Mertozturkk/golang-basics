package main

type Tshirt struct {
	Products
	Size  string
	Color string
}

func (t Tshirt) calculatePrice() int64 {
	var clothingDiscount float64

	clothingDiscount = float64(t.Price) * 0.05
	return t.Price - int64(clothingDiscount)
}
