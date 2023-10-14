package array_slice

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums { // range her iterasyonda indexi ve indexin degerini doner
		sum += num
	}

	/*for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	*/
	return sum
}
