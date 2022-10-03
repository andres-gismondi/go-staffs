package linked

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
	Last *Node

	aux *Node
}

func NewLinkedList(val int) *List {
	n := &Node{Value: val}
	return &List{
		Head: n,
		Last: n,
		aux:  n,
	}
}

func (l *List) Enqueue(val int) {
	node := &Node{Value: val}
	l.Last.Next = node
	l.Last = node
}

func (l *List) EnqueueMiddle(val int, after int) {
	node := l.Head

	newNode := &Node{Value: val}
	var last *Node
	for node.Next != nil {
		last = node
		node = node.Next

		if after == last.Value {
			last.Next = newNode
			newNode.Next = node
			return
		}
	}

	l.Last.Next = newNode
}

func (l *List) EnqueueFront(val int) {
	node := &Node{Value: val, Next: l.Head}
	l.Head = node
}

func (l *List) Dequeue() *Node {
	node := l.Head
	rNode := l.Last

	var last *Node
	for node.Next != nil {
		last = node
		node = node.Next
	}
	last.Next = nil
	l.Last = last
	return rNode
}

func (l *List) DequeueHead() *Node {
	aux := l.Head
	l.Head = l.Head.Next
	return aux
}

func (l *List) Next() *Node {
	res := l.aux
	l.aux = l.aux.Next
	return res
}

func (l *List) ToString() {
	head := l.Head
	for head != nil {
		fmt.Printf("%v ;", head.Value)
		head = head.Next
	}
	fmt.Println()
}
