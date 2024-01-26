//FIRST DAY CHALLENGE

//https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go

//SUM OF POSITIVE

package main

import "fmt"

func main() {
	input := []int{1, -3, 2, 5, -1}
	fmt.Print(positiveSum(input))
}

func positiveSum(numbers []int) int {
	result := 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i] >= 0 {
			result += numbers[i]
		}
	}

	return result
}
