package practices_test

import (
	"fmt"
	"testing"
	"time"

	"go-staffs/practices"
)

func TestEasyProblems_Sort(t *testing.T) {
	numbers := []int64{10, 9, 8, 1, 2, 3, 4, 5}
	res := practices.SortSlice(numbers)

	fmt.Println(res)
}

func Test_Currification(t *testing.T) {
	testFunc := func() func() int {
		a, b := 1, 1
		return func() int {
			fmt.Println(time.Now())
			a++
			b++
			return a + b
		}
	}

	tf := testFunc()
	time.Sleep(time.Second + 1)
	res := tf()
	fmt.Println(res)
	time.Sleep(time.Second + 1)
	res = tf()
	fmt.Println(res)
}
