package baitap

import (
	"fmt"
)

func Slove4() {
	array := []int {3,2,5,1,7}

	minMoves := calculateMinMoves(array)
	fmt.Println(minMoves)
}

func calculateMinMoves(array []int) int {
	minMoves := 0

	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			moves := array[i-1] - array[i]
			array[i] += moves
			minMoves += moves
		}
	}
	fmt.Println(array)
	return minMoves
}