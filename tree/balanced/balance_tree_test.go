package balanced_test

import (
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
