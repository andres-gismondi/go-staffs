package practices

import (
	"fmt"
	"sync"
)

func MergeArrays1(a [][]int) {
	cnnls := make(chan []int, len(a))

	merge := func() {
		defer close(cnnls)

		var total []int
		for i := 0; i < len(a); i++ {
			arr := <-cnnls
			total = append(total, arr...)
		}

		sortSlice(total, func(ints []int) {
			fmt.Println(ints)
		})
	}

	fn := func() func([]int) {
		return func(ints []int) {
			cnnls <- ints
		}
	}
	for _, arr := range a {
		go sortSlice(arr, fn())
	}
	go merge()
}

func MergeArrays2(a [][]int) {
	cnnls := make(chan []int, len(a))
	quit := make(chan bool)

	go merge2(cnnls, quit)

	var wg sync.WaitGroup
	fn := func() func([]int) {
		return func(ints []int) {
			defer wg.Done()
			cnnls <- ints
		}
	}
	for _, arr := range a {
		wg.Add(1)
		go sortSlice(arr, fn())
	}

	wg.Wait()
	quit <- true
}

func merge2(chnnls chan []int, quit chan bool) {
	defer func() {
		close(chnnls)
		close(quit)
	}()

	var total []int
	exit := true
	for exit {
		select {
		case ch := <-chnnls:
			total = append(total, ch...)
		case <-quit:
			exit = false
		}
	}

	sortSlice(total, func(ints []int) {
		fmt.Println(ints)
	})

}

func sortSlice(S []int, fn func([]int)) {
	for i := len(S); i > 0; i-- {
		for j := 1; j < i; j++ {
			if S[j-1] > S[j] {
				intermediate := S[j]
				S[j] = S[j-1]
				S[j-1] = intermediate
			}
		}
	}
	fn(S)
}
