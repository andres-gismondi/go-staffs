package generics

import (
	"fmt"
	"testing"
)

type Number interface {
	int64 | float64
}

func Test_TypeConstraint(t *testing.T) {
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
		SumNumbers(ints),
		SumNumbers(floats))
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func TestArray(t *testing.T) {
	sarasa := []int{1, 2, 3, 4, 5, 6}
	hh := sarasa[:1]
	fmt.Println(hh)
	hh = sarasa[1:]
	fmt.Println(hh)
	hh = sarasa[1:1]
	fmt.Println(hh)
	hh = sarasa[2:]
	fmt.Println(hh)
}
