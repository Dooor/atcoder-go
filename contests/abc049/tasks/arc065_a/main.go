package main

import (
	"fmt"
)

func check(s string, d []string) bool {
	for i := 0; i < len(d); i++ {
		if len(s) < len(d[i]) {
			continue
		}

		if s[0:len(d[i])] == d[i] {
			if s == d[i] {
				return true
			}
			return check(s[len(d[i]):], d)
		}
	}
	return false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var s string
	fmt.Scan(&s)
	s = reverse(s)

	d := []string{"esare", "resare", "maerd", "remaerd"}

	if check(s, d) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
