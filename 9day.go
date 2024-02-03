// 9/366

// https://www.codewars.com/kata/56f69d9f9400f508fb000ba7/train/go

// COUNT THE MONKEYS

package main

import "fmt"

func main() {
	fmt.Print(monkeyCount(10))
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
