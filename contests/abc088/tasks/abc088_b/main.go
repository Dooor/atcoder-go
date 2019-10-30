package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n      int
		result int
	)

	fmt.Scan(&n)

	cards := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&cards[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cards)))

	for i := 0; i < len(cards); i++ {
		if i%2 == 0 {
			result += cards[i]
		} else {
			result -= cards[i]
		}
	}

	fmt.Println(result)
}
