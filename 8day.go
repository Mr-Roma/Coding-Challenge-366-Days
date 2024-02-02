// 8/366

// https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go

// VOWEL COUNT

package main

import "fmt"

func main() {
	fmt.Print(GetCount("localhost"))
}

func GetCount(str string) int {
	var count int

	count = 0

	for i := 0; i < len(str); i++ {
		char := str[i]
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++
		}
	}

	return count
}
