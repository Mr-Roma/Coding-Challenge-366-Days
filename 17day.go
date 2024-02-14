// 17/366

// https://www.codewars.com/kata/515e271a311df0350d00000f/train/go

// Square(n) Sum

package main

import "fmt"

func main() {
	number := []int{1, 2, 3}
	fmt.Print(squareSum(number))
}

func squareSum(numbers []int) int {
	var sum = 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i] * numbers[i]
	}
	return sum
}
