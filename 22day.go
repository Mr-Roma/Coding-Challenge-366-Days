// 22/366

// https://www.codewars.com/kata/5a3fe3dde1ce0e8ed6000097/train/go

// Century From Year

package main

import "fmt"

func main() {
	years := 1501
	fmt.Print(century(years))
}

func century(year int) int {
	var century int

	if year%100 == 0 {
		century = year / 100
	} else {
		century = (year / 100) + 1
	}

	return century
}
