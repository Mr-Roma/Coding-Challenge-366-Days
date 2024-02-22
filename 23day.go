// 23/366

// https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3/train/go

// Abbreviate a two words name

package main

import (
	"fmt"
	"strings"
)

func main() {
	inputString := "Maria Luisa"
	fmt.Print(AbbrevName(inputString))
}

func AbbrevName(name string) string {
	names := strings.Split(name, " ")
	if len(names) != 2 {
		return "Invalid Input, ,the input should separate by koma"
	}

	initials := strings.ToUpper(string(names[0][0])) + "." + strings.ToUpper(string(names[1][0]))

	return initials
}
