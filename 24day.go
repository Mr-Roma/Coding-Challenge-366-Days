// 24/366

// https://www.codewars.com/kata/582cb0224e56e068d800003c/train/go

// Keep Hydrated !

package main

import "fmt"

func main() {
	waktu := 11.8
	fmt.Print(Litres(waktu))
}

func Litres(time float64) int {
	return int(time / 2)
}
