package balanced_test

import (
	"fmt"
	"testing"

	"go-staffs/tree/balanced"
)

func TestTree(t *testing.T) {
	tr := balanced.NewTree(7)
	tr.Add(8)
	tr.Add(1)
	tr.Add(9)

	tr.Print()
}

func Test_testMaster(t *testing.T) {
	fmt.Println("test commit on branch")
}
