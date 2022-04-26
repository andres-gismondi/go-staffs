package function_type

import (
	"fmt"
	"testing"
)

type Testing interface {
	Execute() error
}

type Retrieve func() error

type ImplTest struct{}

func (i ImplTest) Execute() error {
	fmt.Println("executing implementation....")
	return nil
}

func checkCall(fn Retrieve) {
	fmt.Println("checkCall...")
	anotherCall(fn, fn())
}
func anotherCall(fn Retrieve, err error) {
	if err == nil {
		fmt.Println("executed when call it in the parameter")
	}
	fmt.Println("anotherCall is going to execute Testing...")
	fn()
}

func TestFunction(t *testing.T) {
	checkCall(func() error {
		fmt.Println("you can execute wherever functions here... or implementation")
		return nil
	})

	impl := ImplTest{}
	checkCall(impl.Execute)

	fn := func() error {
		fmt.Println("like anonymous functions...")
		return nil
	}
	checkCall(fn)
}
