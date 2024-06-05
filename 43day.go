// 43/366

// https://www.codewars.com/kata/55225023e1be1ec8bc000390/train/go

// Jenny's secret message

package main

import "fmt"

func main() {
	fmt.Println(Greet("Johnny"))
}

func Greet(name string) string {
	if name == "Johnny" {
		return fmt.Sprintf("Hello, my love!")
	} else {
		return fmt.Sprintf("Hello, %v!", name)
	}
}
