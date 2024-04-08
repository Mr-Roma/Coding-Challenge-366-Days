// 37/366

// https://www.codewars.com/kata/5ce9c1000bab0b001134f5af/train/go

// Quarter of the Year

package main

import "fmt"

func main() {
	months := []int{1, 2, 3, 6, 11, 12}

	for _, month := range months {
		quarter := QuarterOf(month)
		if quarter != -1 {
			fmt.Printf("Month %d belongs to the %d quarter\n", month, quarter)
		} else {
			fmt.Printf("Invalid month: %d\n", month)
		}
	}
}

func QuarterOf(month int) int {
	if month < 1 || month > 12 {
		return -1 // Invalid month
	}
	return (month-1)/3 + 1
}
