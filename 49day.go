// 49/366

// https://www.codewars.com/kata/56747fd5cb988479af000028/train/go

// Get the Middle Character

package main

import "fmt"

func main() {
	fmt.Println(GetMiddle("ayo"))
}

func GetMiddle(s string) string {
	n := len(s)
	if n%2 == 0 {
		// If the length of the string is even, return the middle 2 characters
		return s[n/2-1 : n/2+1]
	}
	// If the length of the string is odd, return the middle character
	return string(s[n/2])
}
