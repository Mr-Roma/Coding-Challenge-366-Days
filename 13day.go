// 13/366

// https://www.codewars.com/kata/5bb904724c47249b10000131/train/go

// FIND THE SMALLEST INTEGER IN THE ARRAY

package main

import "fmt"

func main() {
	number := []int{13, 7, 45, 78}
	fmt.Print(smallestIntegerFinder(number))
}

func smallestIntegerFinder(numbers []int) int {
	smallest := 99999
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}
	return smallest
}
