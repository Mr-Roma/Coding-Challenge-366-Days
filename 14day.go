// 14/366

// https://www.codewars.com/kata/53369039d7ab3ac506000467/train/go

// CONVERT BOOLEAN VALUES TO STRING "YES" OR "NO"

package main

import "fmt"

func main() {
	fmt.Print(BoolToWord(true))
}

func BoolToWord(word bool) string {
	result := "No"
	if word == true {
		result = "Yes"
	}
	return result
}
