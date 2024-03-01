package main

import "fmt"

func main() {
	fmt.Print(Past(0, 1, 1))
}

func Past(h, m, s int) int {
	h = (h * 60 * 60) * 1000
	m = (m * 60) * 1000
	s = s * 1000

	total := m + h + s

	return total
}
