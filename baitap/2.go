package baitap

import (
	"fmt"
)

func Slove2() {
	var n int = 5;
	expectedSum := (n * (n + 1)) / 2

	var arr = []int{1,2,5,4}

	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	missingNumber := expectedSum - sum
	fmt.Println(missingNumber)
}