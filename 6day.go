// 6/366

// https://www.codewars.com/kata/5513795bd3fafb56c200049e/train/go

// COUNT BY X

package main

import "fmt"

func main() {
	fmt.Print(countBy(1, 10))
}

func countBy(x, n int) []int {
	array := []int{}
	for i := 1; i <= n; i++ {
		array = append(array, x*i)
	}
	return array
}
