package baitap
// Number Spiral

import "fmt"

func numberSpiral(Y, X int) {
	if Y > X {
		ans := (Y - 1) * (Y - 1)
		var add int
		if Y%2 != 0 {
			add = X
		} else {
			add = 2*Y - X
		}

		fmt.Println(ans + add)
	} else { 
		ans := (X - 1) * (X - 1)
		var add int
		if X%2 == 0 {
			add = Y
		} else {
			add = 2*X - Y
		}
		fmt.Println(ans + add)
	}
}

func Slove6() {
	Y := 4
	X := 2
	numberSpiral(Y, X)
}
