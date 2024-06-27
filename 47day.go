// 47/366

// https://www.codewars.com/kata/52fba66badcd10859f00097e/train/go

// Disemvowel Trolls

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Disemvowel("This website is for losers LOL!"))
}

func Disemvowel(comment string) string {
	vowels := "aeiouAEIOU"
	for _, v := range vowels {
		comment = strings.ReplaceAll(comment, string(v), "")
	}
	return comment
}
