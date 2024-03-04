// 30/366

// https://www.codewars.com/kata/55a70521798b14d4750000a4/train/go

// Returning Strings

package main

import "fmt"

func main() {
	namer := "Anton"
	fmt.Print(Greet(namer))
}

func Greet(name string) string {
	return "Hello" + ", " + name + " " + "how are you doing today?"
}
