package practices_test

import (
	"fmt"
	"testing"
	"time"

	"go-staffs/practices"
)

func TestProblems_MergeArraysWithLocalFunction(t *testing.T) {
	a := [][]int{
		{2, 7, 11, 4},
		{8, 14, 3, 6},
		{12, 5, 9, 16},
		{10, 15, 13, 1},
	}

	practices.MergeArrays1(a)

	time.Sleep(time.Second * 2)
}

func TestProblems_MergeArraysWithExternalFunction(t *testing.T) {
	a := [][]int{
		{2, 7, 11, 4},
		{8, 14, 3, 6},
		{12, 5, 9, 16},
		{10, 15, 13, 1},
	}

	practices.MergeArrays2(a)

	time.Sleep(time.Second * 2)
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

func TestDeadLock(t *testing.T) {
	var stocks map[string]float64 // stock -> price
	price := stocks["MSFT"]
	fmt.Printf("%f\n", price)
}

func TestSlicePointers(t *testing.T) {
	s1 := make([]int, 10, 100)
	s2 := s1

	s1[0] = 00
	s2[1] = 11

	fmt.Println(s1)
	s2[0] = 22
	fmt.Println(s1)
}

func TestSliceInit(t *testing.T) {
	s1 := make([]int, 1)
	fmt.Println(s1)

	var s2 []int
	fmt.Println(s2)
}

func TestContinue(t *testing.T) {
	w := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range w {
		if num == 2 || num == 3 {
			continue
		}
		fmt.Println(num)
	}
}
