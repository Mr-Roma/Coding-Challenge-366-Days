// 34/366

// https://www.codewars.com/kata/576bb71bbbcf0951d5000044/train/go

// Count of Positives/Sum of Negatives

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Print(CountPositvesSumNegatives(arr))
}

func CountPositvesSumNegatives(numbers []int) []int {
	if numbers == nil || len(numbers) == 0 {
		return []int{}
	}
	positiveCount := 0
	negativeSum := 0

	for _, num := range numbers {
		if num > 0 {
			positiveCount += 1
		} else {
			negativeSum += num
		}
	}
	return []int{positiveCount, negativeSum}

}
