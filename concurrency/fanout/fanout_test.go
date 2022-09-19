package fanout_test

import (
	"fmt"
	"testing"

	"go-staffs/concurrency/fanout"
)

func TestFanOut(t *testing.T) {
	faout := fanout.FanOut{
		Values: []int{1, 2, 3, 4, 5, 6},
	}
	faout.Handler()
}

func TestReverseArray(t *testing.T) {
	a := []int32{9, 5, 6, 7}

	var control int32
	for j := 0; j < len(a); j++ {
		for i := 0; i < len(a); i++ {
			if i+1 <= len(a)-1 {
				if a[i] < a[i+1] {
					control = a[i]
					a[i] = a[i+1]
					a[i+1] = control
				}
			}
		}
	}

	fmt.Println(a)
}
