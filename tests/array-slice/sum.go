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

func SumAll(numsToSum ...[]int) []int {
	/*var sums []int
	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}

	*/
	lengthOfNumbers := len(numsToSum)
	sums := make([]int, lengthOfNumbers) // make ile slice olusturulabilir
	for i, nums := range numsToSum {
		sums[i] = Sum(nums)
	}
	return sums
}
