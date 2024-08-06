package initialize

import "fmt"

func TryGeneric() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}

type Number interface {
	int64 | float64
}

// func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
