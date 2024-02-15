// 18/366

// https://www.codewars.com/kata/57a0e5c372292dd76d000d7e/train/go

// String Repeat

package main

import "fmt"

func main() {
	var n int
	var value string

	fmt.Scan(&n, &value)
	fmt.Print(RepeatStr(n, value))
}

func RepeatStr(repetitions int, value string) string {
	var firstVal, totalStr string

	firstVal = value
	for i := 0; i < repetitions; i++ {
		totalStr += firstVal
	}
	return totalStr
}
