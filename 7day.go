// 7/366

// https://www.codewars.com/kata/5556282156230d0e5e000089/train/go

// DNA TO RNA

// FIRST VERSION
package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter DNA sequence: ")
	fmt.Scanln(&input)

	validChars := "GCAT"

	// Check if input contains only valid characters
	if strings.ContainsAny(input, validChars) {
		fmt.Println("String Valid")
		fmt.Println(DNAToRNA(input))
	} else {
		fmt.Println("String not valid")
	}
}

func DNAToRNA(dna string) string {
	// Replace "T" with "U" in the DNA sequence
	runeSlice := []rune(dna)
	for i := 0; i < len(runeSlice); i++ {
		if runeSlice[i] == 'T' {
			runeSlice[i] = 'U'
		}
	}
	return string(runeSlice)
}

// SECOND VERSION
func DnaToRna(dna string) string {
	// Replace "T" with "U" in the DNA sequence
	rna := strings.ReplaceAll(dna, "T", "U")
	return rna
}
