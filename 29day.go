// 29/366

// https://www.codewars.com/kata/55f9bca8ecaa9eac7100004a/train/go

// I love you as little, a lot, passionately, madly or at all

package main

import "fmt"

func main() {
	petals := 14 // Change this to the number of petals you want to test
	fmt.Printf("Last phrase: %s\n", HowMuchILoveYou(petals))
}

func HowMuchILoveYou(i int) string {
	phrases := []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}

	lastIndex := (i - 1) % len(phrases)
	return phrases[lastIndex]
}
