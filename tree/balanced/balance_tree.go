package balanced

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	root  *Node
	nodes []*Node
}

func NewTree(value int) *Tree {
	root := &Node{Value: value}
	nodes := append(make([]*Node, 0), root)
	return &Tree{root: root, nodes: nodes}
}

func (t *Tree) Add(value int) {
	t.addNode(value, t.root)
}

func (t *Tree) addNode(value int, node *Node) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value}
			t.nodes = append(t.nodes, node.Left)
			return
		}
		t.addNode(value, node.Left)
	}
	if value > node.Value {
		if node.Right == nil {
			node.Right = &Node{Value: value}
			t.nodes = append(t.nodes, node.Right)
			return
		}
		t.addNode(value, node.Right)
	}
	return
}

func (t *Tree) Print() {
	t.printNodes(t.root)
	fmt.Printf("Total elements: %v", len(t.nodes))
}

func (t *Tree) printNodes(node *Node) {
	if node == nil {
		return
	}
	t.printNodes(node.Left)
	fmt.Printf("%v ", node.Value)
	t.printNodes(node.Right)
}
