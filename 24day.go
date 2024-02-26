package main

import "fmt"

func main() {
	waktu := 11.8
	fmt.Print(Litres(waktu))
}

func Litres(time float64) int {
	return int(time / 2)
}
