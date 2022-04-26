package list

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
}

func NewList(value int) *List {
	return &List{Head: &Node{Value: value}}
}

func (l *List) Add(value int) {
	if l == nil {
		panic("list must be initialized")
	}

	newNode := &Node{Value: value, Next: l.Head}
	l.Head = newNode
}

func (l *List) AddNode(node *Node) {
	if l == nil {
		panic("list must be initialized")
	}

	node.Next = l.Head
	l.Head = node
}

func (l *List) Next() {
	l.Head = l.Head.Next
}

func (l *List) Print() {
	first := l.Head
	for {
		if l.Head == nil {
			break
		}

		fmt.Printf("%v, ", l.Head.Value)
		l.Head = l.Head.Next
	}
	l.Head = first
}
