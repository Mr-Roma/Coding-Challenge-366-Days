// 5/366

// https://www.codewars.com/kata/55685cd7ad70877c23000102/train/go

// RETURN NEGATIVE

package main

import "fmt"

func main() {
	input := 5
	input2 := []int{4, 16, 12, 14, 8, 16}
	fmt.Println(makeNegative(input))
	fmt.Print(Maps(input2))

}

func makeNegative(x int) int {
	var neg int
	neg = x * (-1)
	if x < 0 {
		x = x * (-1)
	}
	return neg
}

// LOST WITHOUT A MAP

func Maps(x []int) []int {
	result := make([]int, len(x))

	for i := range x {
		result[i] = x[i] << 1
	}

	return result
}
