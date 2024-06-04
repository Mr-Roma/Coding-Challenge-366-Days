// 42/366

// https://www.codewars.com/kata/554b4ac871d6813a03000035/train/go

// Highest and Lowest

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//var array_name = [length]datatype{values}
	var arrNumber = [5]string{"1", "2", "3", "4", "5"}
	strNumber := strings.Join(arrNumber[:], " ")
	fmt.Println(HighAndLow(strNumber))
}

func HighAndLow(i string) string {
	// Split the input string into a slice of string numbers
	stringNumbers := strings.Fields(i)

	// Initialize slices to store integer values
	var numbers []int

	// Iterate over the string numbers
	for i := range stringNumbers {
		// Convert each string number to an integer
		num, err := strconv.Atoi(stringNumbers[i])

		//handle for a production error
		if err != nil {
			panic(err)
		}
		//stored the convert numbers in the numbers variable
		numbers = append(numbers, num)
	}
	//set the highest and lowest in the first number in the array numbers
	highest := numbers[0]
	lowest := numbers[0]

	//looping through the numbers array to get the highest and lowest number
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > highest {
			highest = numbers[i]
		}

		if numbers[i] < lowest {
			lowest = numbers[i]
		}
	}
	//return the highest and lowest number, by convert them into strings
	return fmt.Sprintf("%d %d", highest, lowest)
}
