package main

import (
	"fmt"
	"strconv"
)

// Point is struct
type Point struct {
	row    int
	column int
}

func (point Point) key() string {
	return strconv.Itoa(point.row) + "." + strconv.Itoa(point.column)
}

// Lines is []Line
type Lines map[string]Point

func step(point Point, lines Lines) int {
	next := Point{
		row:    point.row,
		column: point.column - 1,
	}

	if next.column == 0 {
		return next.row
	}

	if swap, ok := lines[next.key()]; ok {
		return step(swap, lines)
	}

	return step(next, lines)
}

func main() {
	var l, n, m int

	fmt.Scan(&l, &n, &m)

	lines := make(Lines, m)

	for i := 0; i < m; i++ {
		var ai, bi, ci int
		fmt.Scan(&ai, &bi, &ci)

		left := Point{
			row:    ai,
			column: bi,
		}
		right := Point{
			row:    ai + 1,
			column: ci,
		}

		lines[left.key()] = right
		lines[right.key()] = left
	}

	start := Point{
		row:    1,
		column: l,
	}

	result := step(start, lines)

	fmt.Println(result)
}
