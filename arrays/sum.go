package main

func main() {

}

func Sum(nums [5]int) (sum int) {
	for i := 0; i < 5; i++ {
		sum += nums[i]
	}
	return sum
}
