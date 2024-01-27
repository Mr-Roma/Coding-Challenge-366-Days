// 3/366

// https://www.codewars.com/kata/53da3dbb4a5168369a0000fe/train/go

// EVEN OR ODD
package main

import "fmt"

func main() {
	number := 5
	fmt.Print(EvenOrOdd(number))
}

func EvenOrOdd(number int) string {
	var evenOdd string

	if number%2 == 0 {
		evenOdd = "Even"
	} else {
		evenOdd = "Odd"
	}
	return evenOdd
}
