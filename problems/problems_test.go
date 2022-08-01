package problems_test

import (
	"testing"
	"time"

	"go-staffs/problems"
)

func TestProblems_MergeArraysWithLocalFunction(t *testing.T) {
	a := [][]int{
		{2, 7, 11, 4},
		{8, 14, 3, 6},
		{12, 5, 9, 16},
		{10, 15, 13, 1},
	}

	problems.MergeArrays1(a)

	time.Sleep(time.Second * 2)
}

func TestProblems_MergeArraysWithExternalFunction(t *testing.T) {
	a := [][]int{
		{2, 7, 11, 4},
		{8, 14, 3, 6},
		{12, 5, 9, 16},
		{10, 15, 13, 1},
	}

	problems.MergeArrays2(a)

	time.Sleep(time.Second * 2)
}
