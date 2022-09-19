package list_merge_points

import (
	"fmt"

	"go-staffs/data_structures/linked_list/list"
)

type Merge struct {
	First  *list.List
	Second *list.List
}

func (m Merge) Merge() {
	exit := false
	sndHead := m.Second
	for !exit {
		if m.First.Head == nil {
			break
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
}
