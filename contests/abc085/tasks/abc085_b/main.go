package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n         int
		lastMochi int
		result    int
	)

	fmt.Scan(&n)

	mochis := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&mochis[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(mochis)))

	for i := 0; i < len(mochis); i++ {
		if lastMochi == 0 || lastMochi > mochis[i] {
			result++
			lastMochi = mochis[i]
		}
	}

	fmt.Println(result)
}
