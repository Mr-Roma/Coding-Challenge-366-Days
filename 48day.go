// 48/366

// https://www.codewars.com/kata/55cb632c1a5d7b3ad0000145/train/go

// Keep up the hoop

package main

import "fmt"

func HoopCount(n int) string {
	message := ""

	if n >= 10 {
		message = "Great, now move on to tricks"
	} else {
		message = "Keep at it until you get it"
	}

	return message
}

func main() {
	var n int
	fmt.Print("Please enter the amount of hoops: ")
	fmt.Scan(&n)

	fmt.Println(HoopCount(n))
}
