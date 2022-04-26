package generics_test

import (
	"fmt"
	"testing"
)

func Test_SimpleSum(t *testing.T) {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		GenericSum(ints),
		GenericSum(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func GenericSum[L comparable, V int64 | float64](m map[L]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
