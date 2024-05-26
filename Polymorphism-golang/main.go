package main

type purchasable interface {
	calculatePrice() int64
}

var chart []purchasable

func addChart(products ...purchasable) {
	chart = append(chart, products...)
}

func calculateTotal() int64 {
	var total int64
	for _, product := range chart {
		total += product.calculatePrice()
	}
	return total
}

func main() {
	tshirt := Tshirt{Products{Price: 700, Brand: "Nike"}, "M", "Black"}
	monitor := Monitor{Products{Price: 5000, Brand: "Samsung"}, "1920x1080", "24 inch"}
	wine := Wine{Products{Price: 1500, Brand: "Chateau Margaux"}, "Red", 2019}

	addChart(tshirt, monitor, wine)

	total := calculateTotal()
	println(total)

}
