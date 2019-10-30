package main

import (
	"fmt"
	"math"
)

func isMovable(current []int, next []int, t int) bool {
	diffX := next[0] - current[0]
	diffY := next[1] - current[1]

	distance := int(math.Abs(float64(diffX)) + math.Abs(float64(diffY)))

	if t >= distance && ((t%2 != 0 && distance%2 != 0) || (t%2 == 0 && distance%2 == 0)) {
		return true
	}

	return false
}

func main() {
	var (
		n           int
		x           int = 0
		y           int = 0
		currentTime int = 0
	)

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var t, xi, yi int
		fmt.Scan(&t, &xi, &yi)

		if isMovable([]int{x, y}, []int{xi, yi}, t-currentTime) {
			x = xi
			y = yi
			currentTime = t
			continue
		}
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
