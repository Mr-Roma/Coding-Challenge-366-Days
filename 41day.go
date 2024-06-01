// 41/366

// https://www.codewars.com/kata/5b077ebdaf15be5c7f000077/train/go

// If you can't sleep, just count sheep!!

package main

import (
	"fmt"
	"strconv"
)

func countSheep(num int) string {
	var total string
	total = ""
	for i := 1; i <= num; i++ {
		total += strconv.Itoa(i) + " sheep..."
	}
	return total
}

func main() {
	num := 3
	fmt.Println(countSheep(num))
}
