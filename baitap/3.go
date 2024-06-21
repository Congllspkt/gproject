package baitap

import (
	"fmt"
)

func Slove3() {
	longestRepetition := findLongestRepetition("AABBCCC")
	fmt.Println(longestRepetition)
}

func findLongestRepetition(sequence string) int {
	maxLength := 0
	currentLength := 1

	for i := 1; i < len(sequence); i++ {
		if sequence[i] == sequence[i-1] {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1
		}
	}
	if currentLength > maxLength {
		maxLength = currentLength
	}

	return maxLength
}