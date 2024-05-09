// 39/366

// https://www.codewars.com/kata/59ca8246d751df55cc00014c/train/go

// Is he gonna survive?

package main

import "fmt"

func main() {
	fmt.Print(Hero(8, 4))
}

func Hero(bullets, dragons int) bool {
	var total_calc int
	total_calc = dragons * 2

	if total_calc == bullets || total_calc < bullets {
		return true
	}
	return false
}
