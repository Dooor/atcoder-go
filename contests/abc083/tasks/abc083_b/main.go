package main

import (
	"fmt"
)

func sumEveryDiget(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	var (
		n, a, b, result int
	)
	fmt.Scan(&n, &a, &b)

	for i := 0; i <= n; i++ {
		if sum := sumEveryDiget(i); sum >= a && sum <= b {
			result += i
		}
	}
	fmt.Println(result)
}
