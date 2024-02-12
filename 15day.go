// 15/366

// https://www.codewars.com/kata/56dec885c54a926dcd001095/train/go

// OPPOSITE NUMBER

package main

import "fmt"

func main() {
	number := -5
	fmt.Print(Opposite(number))
}

func Opposite(value int) int {
	return value * (-1)
}
