package baitap

import (
	"fmt"
)

func Slove5() {
	var n int = 5

	permutation := constructBeautifulPermutation(n)
	fmt.Print(permutation)
}

func constructBeautifulPermutation(n int) []int {
	if n == 1 {
		return []int{1}
	}

	if n < 4 {
		return nil
	}

	permutation := make([]int, n)
	i := 0

	// Place even numbers in ascending order
	for num := 2; num <= n; num += 2 {
		permutation[i] = num
		i++
	}

	// Place odd numbers in ascending order
	for num := 1; num <= n; num += 2 {
		permutation[i] = num
		i++
	}

	return permutation
}