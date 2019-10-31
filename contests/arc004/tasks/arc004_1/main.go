package main

import (
	"fmt"
	"math"
)

// Position is struct
type Position struct {
	x int
	y int
}

// Positions is []Position
type Positions []Position

func maxDistance(position Position, positions Positions) float64 {
	var max float64

	for i := 0; i < len(positions); i++ {
		diffX := position.x - positions[i].x
		diffY := position.y - positions[i].y
		if result := math.Sqrt(float64(diffX*diffX + diffY*diffY)); result > max {
			max = result
		}
	}

	return max
}

func main() {
	var (
		n   int
		max float64
	)
	fmt.Scan(&n)

	positions := make(Positions, n)

	for i := 0; i < n; i++ {
		var x, y int

		fmt.Scan(&x, &y)

		positions[i] = Position{
			x: x,
			y: y,
		}
	}

	for i := 0; i < n; i++ {
		if result := maxDistance(positions[i], positions[i+1:]); result > max {
			max = result
		}
	}
	fmt.Println(max)
}
