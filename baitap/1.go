package main

import "fmt"

func main() {
	var n int = 3
	sequence := generateSequence(n)
	for _, num := range sequence {
		fmt.Printf("%d ", num)
	}
}

func generateSequence(n int) []int {
	sequence := []int{n}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		sequence = append(sequence, n)
	}

	return sequence
}