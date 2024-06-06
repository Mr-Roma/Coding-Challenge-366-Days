// 44/366

// https://www.codewars.com/kata/55cbd4ba903825f7970000f5/train/go

// Grasshopper - Grade book

package main

import "fmt"

func main() {
	fmt.Println(string(GetGrade(95, 90, 93)))  // A
	fmt.Println(string(GetGrade(70, 70, 100))) // B
	fmt.Println(string(GetGrade(70, 70, 70)))  // C
	fmt.Println(string(GetGrade(65, 70, 59)))  // D
	fmt.Println(string(GetGrade(44, 55, 52)))  // F
}

func GetGrade(a, b, c int) rune {
	avg := int((float64(a) + float64(b) + float64(c)) / 3.0)
	switch {
	case avg >= 90:
		return 'A'
	case avg >= 80:
		return 'B'
	case avg >= 70:
		return 'C'
	case avg >= 60:
		return 'D'
	default:
		return 'F'
	}
}
