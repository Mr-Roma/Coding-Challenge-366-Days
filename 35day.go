// 35/366

// https://www.codewars.com/kata/5a00e05cc374cb34d100000d/train/go

// Reversed Sequence

package main

import "fmt"

func main() {
	n := 5
	fmt.Print(ReverseSeq(n))
}

func ReverseSeq(n int) []int {
	temp := make([]int, n)
	for i := range temp {
		temp[i] = n
		n--
	}
	return temp
}
