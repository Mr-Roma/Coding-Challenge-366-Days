// 16/366

// https://www.codewars.com/kata/515e188a311df01cba000003/train/go

// Get Planet Name By ID

package main

import "fmt"

func main() {
	fmt.Print(GetPlanetName(3))
}

func GetPlanetName(ID int) string {
	switch ID {
	case 1:
		return "Mercury"
	case 2:
		return "Venus"
	case 3:
		return "Earth"
	case 4:
		return "Mars"
	case 5:
		return "Jupiter"
	case 6:
		return "Saturn"
	case 7:
		return "Uranus"
	case 8:
		return "Neptune"
	default:
		return "Pluto"
	}
}
