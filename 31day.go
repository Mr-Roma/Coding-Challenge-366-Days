// 31/366

// https://www.codewars.com/kata/57f780909f7e8e3183000078/train/go

// Beginner Reduce but grow

package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4}
	fmt.Print(Grow(number))
}

func Grow(arr []int) int {
	total := 1
	for i := 0; i < len(arr); i++ {
		total = arr[i] * total
	}
	return total
}
