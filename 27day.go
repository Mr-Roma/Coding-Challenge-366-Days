// 27/366

// https://www.codewars.com/kata/5583090cbe83f4fd8c000051/train/go

// Convert number to reversed array of digits

package main

import "fmt"

func main() {
	num := 35231
	fmt.Print(Digitize(num))
}

func Digitize(n int) []int {
	var result []int

	if n == 0 {
		return []int{0}
	}
	for n > 0 {
		digit := n % 10
		result = append(result, digit)
		digit /= 10
	}
	return result
}
