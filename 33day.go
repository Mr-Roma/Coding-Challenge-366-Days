// 33/366

// https://www.codewars.com/kata/544675c6f971f7399a000e79/train/go

// Convert String to Number

package main

import "fmt"

func main() {
	fmt.Print(StringToNumber("1234"))
}

func StringToNumber(str string) int {
	var result int
	var isMinus bool

	for i := 0; i < len(str); i++ {
		char := str[i]
		if char == '-' {
			isMinus = true
			continue
		}
		digit := int(char - '0')
		result = result*10 + digit
	}
	if isMinus {
		result *= -1
	}
	return result
}
