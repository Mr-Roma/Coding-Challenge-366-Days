// 21/366

// https://www.codewars.com/kata/54edbc7200b811e956000556/train/go

// Counting sheep...

package main

import "fmt"

func main() {
	number := []bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}

	fmt.Print(CountSheeps(number))
}

func CountSheeps(numbers []bool) int {
	totalSheeps := 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == true {
			totalSheeps++
		} else if numbers[i] == false {
			totalSheeps = totalSheeps
		} else {
			fmt.Print("null/undefined")
		}
	}

	return totalSheeps
}
