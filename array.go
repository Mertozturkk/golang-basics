package main

import "fmt"

var array_1 [3]int
var array_2 = [3]int{1, 2, 3}

func main() {
	array_3 := make([]int, 3)
	array_3[0] = 1

	fmt.Println(array_1, array_2, array_3)
	fmt.Printf("array_1 len: %d \n", len(array_1))
	fmt.Printf("array_2 len: %d \n", len(array_2))
}
