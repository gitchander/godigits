package main

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Dice

func main() {
	const min = 9856
	digits := make([]int, 6)
	for i := range digits {
		digits[i] = i + 1
	}
	m := make(map[int]rune)
	for i, digit := range digits {
		m[digit] = rune(min + i)
	}
	for _, digit := range digits {
		fmt.Printf("%d: %c\n", digit, m[digit])
	}
}
