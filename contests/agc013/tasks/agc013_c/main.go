package main

import "fmt"

// Ants is []Ant
type Ants []Ant

// Ant is struct
type Ant struct {
	direction int
	position  float64
}

func main() {
	var n, l, t int
	fmt.Scan(&n, &l, &t)

	ants := make(Ants, n)

	for i := 0; i < n; i++ {
		var x, w int
		fmt.Scan(&x, &w)
		ants[i] = Ant{
			direction: w,
			position:  float64(x),
		}
	}

	fmt.Println(ants)
}
