// 40/366

// https://www.codewars.com/kata/57eaeb9578748ff92a000009/train/go

// Sum Mixed Array?

package main

import "strconv"

func SumMix(arr []any) int {
	total := 0

	for i := 0; i < len(arr); i++ {
		// Check if the element at index i is a string
		if str, ok := arr[i].(string); ok {
			// If it's a string, convert it to an integer
			if num, err := strconv.Atoi(str); err == nil {
				total += num
			}
		} else if num, ok := arr[i].(int); ok {
			// If it's already an integer, simply add it to the total
			total += num
		}
	}

	return total
}
