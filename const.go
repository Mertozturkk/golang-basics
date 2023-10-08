package main

const sabit_1 = "sabit 1"
const (
	sabit_2 = "sabit 2"
	sabit_3 = "sabit 3"
)

const (
	sabit_4 = iota + 4
	sabit_5
	sabit_6
)

func main() {
	println(sabit_1, sabit_2, sabit_3, sabit_4, sabit_5, sabit_6)
}
