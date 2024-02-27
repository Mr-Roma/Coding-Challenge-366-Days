// 26/366

// https://www.codewars.com/kata/5545f109004975ea66000086/train/go

// Is n divisible by x and y

package main

import "fmt"

func main() {
	fmt.Print(IsDivisible(12, 2, 5))
}

func IsDivisible(x, n, y int) bool {
	if (x%n == 0) && (x%y == 0) {
		return true
	}
	return false
}
