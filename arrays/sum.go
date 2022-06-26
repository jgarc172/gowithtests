package integers

// Sum returns the sum of 5 integers
func Sum(nums [5]int) (sum int) {
	for i := 0; i < 5; i++ {
		sum += nums[i]
	}
	return sum
}

// SumSlice returns the sum of nums slice
func SumSlice(nums []int) (sum int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	return sum
}

// SumVar returns the sums for each slice in numSlices
// if numsSlices is nil, sums is nil
func SumVar(numsSlices ...[]int) (sums []int) {
	for _, nums := range numsSlices {
		sum := SumSlice(nums)
		sums = append(sums, sum)
	}
	return sums
}
