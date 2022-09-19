package graph

type Node struct {
	ID         int
	Value      int
	Neighbours []Node
}

func NewNode(id int, val int) Node {
	return Node{
		ID:         id,
		Value:      val,
		Neighbours: make([]Node, 0),
	}
}
func (n *Node) AddNeighbour(node Node) {
	n.Neighbours = append(n.Neighbours, node)
}

type Queue []Node

func (q *Queue) Add(n Node) {
	if exist := q.Contains(n); !exist {
		*q = append(*q, n)
	}
}
func (q *Queue) Contains(n Node) bool {
	for _, nn := range *q {
		if nn.ID == n.ID {
			return true
		}
	}
	return false
}
func (q *Queue) First() Node {
	last := (*q)[0:]
	if len(*q) > 1 {
		*q = (*q)[1:len(*q)]
	} else {
		*q = Queue{}
	}
	return last[0]
}
func (q *Queue) Empty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

type visited struct {
	NodeID int
	Value  bool
}

type visits []visited

func (v visits) WasVisited(id int) bool {
	for _, vv := range v {
		if vv.NodeID == id {
			return true
		}
	}
	return false
}
