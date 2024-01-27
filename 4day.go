// 4/366

// https://www.codewars.com/kata/5708f682c69b48047b000e07/train/python

// MULTIPLY THE NUMBER

package main

import "fmt"

func main() {
	var number int
	number = 3
	result := multiply(number)
	fmt.Print(result)
}

func multiply(number int) int {
	var totalDigit int
	tempNumber := number // Save the original number for calculation

	for tempNumber > 0 {
		tempNumber /= 10
		totalDigit++
	}

	result := number * (5 * totalDigit)
	return result
}
