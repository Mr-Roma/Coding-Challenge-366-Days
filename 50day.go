// 50/366

// https://www.codewars.com/kata/542c0f198e077084c0000c2e/train/go

// Count the divisors of a number

package main

import "fmt"

func main() {
	fmt.Println(Divisors(5))
}

func Divisors(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			total++
		}
	}
	return total
}
