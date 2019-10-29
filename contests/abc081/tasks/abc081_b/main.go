package main

import "fmt"

func checkAllEven(nums []int) bool {
	for i := range nums {
		if nums[i]%2 != 0 {
			return false
		}
	}

	return true
}

func main() {
	var (
		n     int
		count int
	)
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	for checkAllEven(nums) {
		count++
		for i := 0; i < len(nums); i++ {
			nums[i] = nums[i] / 2
		}
	}
	fmt.Println(count)
}
