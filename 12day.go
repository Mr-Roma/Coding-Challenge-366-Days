package main

import (
	"fmt"
	"strings"
)

func main() {
	var Games = []string{"3:1", "2:2", "0:1"}
	fmt.Println(Ponints(Games))
}

func Ponints(games []string) int {
	var result int = 0

	for i := 0; i < len(games); i++ {
		split := strings.Split(games[i], ":")
		x := split[0]
		y := split[1]
		if x == y {
			result += 1
		} else if x > y {
			result += 3
		}
	}

	return result
}
