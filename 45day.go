// 45/366

// https://www.codewars.com/kata/56f69d9f9400f508fb000ba7/train/go

// Count the Monkeys!

package main

import "fmt"

func main() {
	fmt.Println(monkeyCount(5)) // [1, 2, 3, 4, 5]
}

func monkeyCount(n int) []int {
	var total []int
	if n > 0 {
		for i := 1; i <= n; i++ {
			total = append(total, i) // Append elements to the slice
		}
	}
	return total
}
