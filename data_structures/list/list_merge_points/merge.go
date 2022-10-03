package list_merge_points

import (
	"fmt"

	"go-staffs/data_structures/list/linked"
)

type Merge struct {
	First  *linked.List
	Second *linked.List
}

func (m Merge) Merge() Info {
	exit := false
	sndHead := m.Second
	for !exit {
		if m.First.Head == nil {
			return mergeInfo("asdsa")
		}
		for {
			if m.Second.Head == nil {
				break
			}
			if m.First.Head == m.Second.Head {
				fmt.Printf("Find: %v", m.First.Head.Value)
				exit = true
			}
			m.Second.Next()
		}

		m.First.Next()
		m.Second = sndHead
	}

	return nil
}

type Info interface {
	Message() string
}

type mergeInfo string

func (mi mergeInfo) Message() string {
	return string(mi)
}
