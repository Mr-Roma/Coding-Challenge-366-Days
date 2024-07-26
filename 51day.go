// 51/366

// https://www.codewars.com/kata/56269eb78ad2e4ced1000013/train/go

// Find the next perfect square!

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindNextSquare(121))
	fmt.Println(FindNextSquare(625))
	fmt.Println(FindNextSquare(114))

}

func FindNextSquare(sq int64) int64 {
	root := int64(math.Sqrt(float64(sq)))

	// Check if the input number is a perfect square
	if root*root == sq {
		// Return the next perfect square
		return (root + 1) * (root + 1)
	} else {
		// Return -1 if the input number is not a perfect square
		return -1
	}
}
