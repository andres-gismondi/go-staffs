package practices

func FibonacciIterative(value int) {
	a, b := 0, 1
	for i := 0; i < value; i++ {
		a, b = b, a+b
	}
}

func FibonacciNormalRecursive(actualRes, beforeRest, base int) int {
	if base < 1 {
		return actualRes
	}
	//fmt.Println(actualRes)
	return FibonacciNormalRecursive(actualRes+beforeRest, actualRes, base-1)
}

func FibonacciDoubleRecursive(base int) {
	//res := 0
	for i := 0; i < base; i++ {
		fibonacciDoubleRecursive(i)
		//fmt.Println(res)
	}
}
func fibonacciDoubleRecursive(value int) int {
	if value <= 1 {
		return 1
	}
	return fibonacciDoubleRecursive(value-1) + fibonacciDoubleRecursive(value-2)
}

func FibonacciCurrification(base int) {
	fib := func() func() int {
		a, b := 0, 1
		return func() int {
			a, b = b, a+b
			return a
		}
	}
	f := fib()
	for i := 0; i < base; i++ {
		//fmt.Println(f())
		f()
	}
}
func fibonacciCurrification() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
