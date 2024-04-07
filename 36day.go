//# 36/366

// https://www.codewars.com/kata/576bb71bbbcf0951d5000044/train/go

// Count Positives/Sum Negatives

package main

import "fmt"

func main() {
	fmt.Println(CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
}

func CountPositivesSumNegatives(numbers []int) []int {
	var save []int
	var positive int
	var sumNegative int

	positive = 0
	sumNegative = 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i] >= 0 {
			positive++
		} else {
			sumNegative += numbers[i]
		}
	}
	save = append(save, positive)
	save = append(save, sumNegative)
	return save
}
