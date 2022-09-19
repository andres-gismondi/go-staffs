package list_test

import (
	"fmt"
	"testing"

	"go-staffs/data_structures/linked_list/list"
)

func TestList(t *testing.T) {
	ll := list.NewList(10)
	ll.Add(11)
	ll.Add(12)
	ll.Add(13)

	ll.Print()
	fmt.Println()
	ll.Next()
	ll.Print()
	fmt.Println()
	ll.Next()
	ll.Print()
	fmt.Println()
}
