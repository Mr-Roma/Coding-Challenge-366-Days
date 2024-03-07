// 32/366

// https://www.codewars.com/kata/555086d53eac039a2a000083/train/go

// Opposite Attract

package main

import "fmt"

func main() {
	var flower1, flower2 int
	flower1 = 8
	flower2 = 1
	fmt.Print(LoveFunc(flower1, flower2))
}

func LoveFunc(flower1, flower2 int) bool {
	if (flower1%2 == 0 && flower2%2 != 0) || (flower1%2 != 0 && flower2%2 == 0) {
		return true
	}
	return false
}
