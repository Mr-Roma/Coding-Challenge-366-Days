// 46/366

// https://www.codewars.com/kata/59342039eb450e39970000a6/train/go

// Count Odd Numbers below n!

package main

import "fmt"

func main() {
	fmt.Println(OddCount(15))
}

func OddCount(n int) int {
	// count := 0
	// n--
	// for n >= 0 {
	// 	if n%2 != 0 {
	// 		count++
	// 	}
	// 	n--
	// }
	// return count

	return n / 2
}
