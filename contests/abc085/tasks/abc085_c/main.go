package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		n      int
		y      int
		result string = "-1 -1 -1"
	)

	fmt.Scan(&n, &y)

	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			k := n - i - j

			if y == 10000*i+5000*j+1000*k {
				result = strconv.Itoa(i) + " " + strconv.Itoa(j) + " " + strconv.Itoa(k)
			}
		}
	}

	fmt.Println(result)
}
