package practices_test

import (
	"fmt"
	"testing"

	"go-staffs/practices"
)

func TestEasyProblems_Sort(t *testing.T) {
	numbers := []int64{10, 9, 8, 1, 2, 3, 4, 5}
	res := practices.SortSlice(numbers)

	fmt.Println(res)
}
