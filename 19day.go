// 19/366

// https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad/train/go

// Invert Values

package main

import "fmt"

func main() {
	num := []int{1, -2, 3, 4}
	fmt.Print(Invert(num))

}

func Invert(arr []int) []int {
	inverted := make([]int, len(arr))

	for _, num := range arr {
		inverted = append(inverted, -num)
	}
	return inverted[len(arr):]
}
