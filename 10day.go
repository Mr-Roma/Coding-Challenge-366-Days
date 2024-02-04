// 10/366

// hhttps://www.codewars.com/kata/5265326f5fda8eb1160004c8/train/go

// CONVERT A NUMBER TO A STRING

package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 123
	fmt.Println(number)
}

func NumberToString(n int) string {
	return strconv.Itoa(n)
}
