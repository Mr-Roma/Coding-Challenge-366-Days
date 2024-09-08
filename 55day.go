// 55/366

// https://www.codewars.com/kata/55f2b110f61eb01779000053/train/go

// Beginner Series #3 Sum of Numbers

package main

func main() {

}

func GetSum(a, b int) int {

	if a == b {
		return a
	}

	if a > b {
		a, b = b, a
	}

	total := 0
	for i := a; i <= b; i++ {
		total += i
	}
	return total
}
