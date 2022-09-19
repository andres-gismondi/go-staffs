package fanout

import (
	"fmt"
	"time"
)

type FanOut struct {
	Values []int
}

func (fo FanOut) Handler() {
	chanIn := make(chan int, len(fo.Values))
	chanOutA := make(chan int, 1)
	chanOutB := make(chan int, 1)

	go fo.fillInChanIn(&chanIn)
	go fo.fanOut(chanIn, chanOutA, chanOutB)
	/*
		for item := range chanOutA {
			fmt.Println(item)
		}
		for item := range chanOutB {
			fmt.Println(item)
		}*/
	time.Sleep(time.Second * 5)
}

func (fo FanOut) fillInChanIn(chanIn *chan int) {
	for item := range fo.Values {
		*chanIn <- item
	}
	close(*chanIn)
}

func (fo FanOut) fanOut(chanIn <-chan int, chanOutA, chanOutB chan int) {
	for item := range chanIn {
		select {
		case chanOutA <- item:
			fmt.Printf("chan A: %v\n", <-chanOutA)
		case chanOutB <- item:
			fmt.Printf("chan B: %v\n", <-chanOutB)
		}
	}
}
