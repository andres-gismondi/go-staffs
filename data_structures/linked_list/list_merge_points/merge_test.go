package list_merge_points_test

import (
	"testing"

	"go-staffs/data_structures/linked_list/list"
	"go-staffs/data_structures/linked_list/list_merge_points"
)

func TestMerge(t *testing.T) {
	lf := list.NewList(10)
	lf.Add(11)
	lf.Add(12)
	lf.Add(13)
	node := list.Node{Value: 14}
	lf.AddNode(&node)

	ls := list.NewList(13)
	ls.AddNode(&node)
	ls.Add(15)
	ls.Add(16)

	m := list_merge_points.Merge{First: lf, Second: ls}
	m.Merge()
}
