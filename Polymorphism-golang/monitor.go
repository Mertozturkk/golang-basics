package main

type Monitor struct {
	Products
	Resolution string
	Size       string
}

func (m Monitor) calculatePrice() int64 {
	var electronicTax float64

	electronicTax = float64(m.Price) * 0.16
	return m.Price + int64(electronicTax)
}
