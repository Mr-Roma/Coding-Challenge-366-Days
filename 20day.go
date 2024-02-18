// 20/366

// https://www.codewars.com/kata/55d24f55d7dd296eb9000030/train/go

// Grasshoper Summation

package main

import (
	"fmt"
)

func main() {
	number := 2
	fmt.Print(Summation(number))
}

func Summation(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}

	return result
}
