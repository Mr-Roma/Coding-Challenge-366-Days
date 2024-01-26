//FIRST DAY CHALLENGE

//https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go

//REVERSE STRING

package main

import "fmt"

func Solution(word string) string {
	r := []rune(word)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	input := "Hello, World!"
	reversed := Solution(input)
	fmt.Println(reversed)
}
