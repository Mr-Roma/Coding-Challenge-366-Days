// 53/366

// https://www.codewars.com/kata/5656b6906de340bd1b0000ac/train/go

// Two to One

package main

import "fmt"

func main() {
	a := "xyaabbbccccdefww"
	b := "xxxxyyyyabklmopq"

	a = "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(TwoToOne(a, b))
}

func TwoToOne(s1 string, s2 string) string {
	s := s1 + s2

	m := make(map[rune]bool)

	for i := 0; i < len(s); i++ {
		m[rune(s[i])] = true
	}

	var uniqueChars []rune

	for r := range m {
		uniqueChars = append(uniqueChars, r)
	}

	for i := 0; i < len(uniqueChars); i++ {
		for j := 0; j < len(uniqueChars)-i-1; j++ {
			if uniqueChars[i] > uniqueChars[j+1] {
				uniqueChars[i], uniqueChars[j+1] = uniqueChars[j+1], uniqueChars[i]
			}
		}
	}

	result := ""

	for _, r := range uniqueChars {
		result += string(r)
	}

	return result
}
