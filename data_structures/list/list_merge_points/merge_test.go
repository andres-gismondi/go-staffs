package list_merge_points_test

import (
	"testing"

	"go-staffs/data_structures/list/linked"
)

func TestMerge(t *testing.T) {
	lf := linked.NewLinkedList(10)
	lf.Enqueue(11)
	lf.Enqueue(12)
	lf.Enqueue(13)

	//node := linked.Node{Value: 14}
	//lf.AddNode(&node)

	ls := linked.NewLinkedList(13)
	//ls.AddNode(&node)
	ls.Enqueue(15)
	ls.Enqueue(16)

	//m := list_merge_points.Merge{First: lf, Second: ls}
	//m.Merge()
}
