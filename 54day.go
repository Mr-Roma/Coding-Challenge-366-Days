// 54/366

// https://www.codewars.com/kata/5390bac347d09b7da40006f6/train/go

// Jaden Casing Strings

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "How can mirrors be real if our eyes aren't real"
	fmt.Println(ToJadenCase(str))
}

func ToJadenCase(str string) string {
	words := strings.Fields(str)
	for i, word := range words {
		if len(word) > 1 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		} else {
			words[i] = strings.ToUpper(word)
		}
	}
	return strings.Join(words, " ")
}
