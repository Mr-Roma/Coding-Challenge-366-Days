package main

import "fmt"

func main() {
	var P1, P2 string
	fmt.Scanln(&P1, &P2)
	fmt.Print(Rps(P1, P2))
}

func Rps(p1, p2 string) string {
	var result string
	if (p1 == "Paper" && p2 == "Rock") || (p1 == "Rock" && p2 == "Scissor") || (p1 == "Scissor" && p2 == "Paper") {
		result = "Player 1 won"
	} else if p1 == p2 {
		result = "Draw"
	} else {
		result = "Player 2 won"
	}
	return result
}
