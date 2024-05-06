// 38/366

// https://www.codewars.com/kata/57eae65a4321032ce000002d/train/go

// Fake Binary

package main

import "fmt"

func main() {
	input := "1234567890"
	output := FakeBin(input)
	fmt.Println(output) // Output should be "0000011111"
}

func FakeBin(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		digit := s[i] - '0'
		if digit < 5 {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}
