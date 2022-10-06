package parser

import (
	"strings"
)

// Money approach is used in order to not use float64, since this kind of operations
// generate bugs in floating points. So the base money is an integer
type Money int64

const (
	USD  Money = 100000000
	GOLD       = USD * 1750
)

// NetPos is the net position
type NetPos struct {
	Gold float64
	USD  Money
}

// AddGold add gold quantity and also the operation to subtract usd
func (np *NetPos) AddGold(q float64, total int64) {
	np.Gold += q

	aux := float64(GOLD) * q
	np.USD = np.USD - Money(aux)
}

// SellGold subtract gold from net position and also sum usd
func (np *NetPos) SellGold(q float64, total int64) {
	np.Gold -= q

	aux := float64(GOLD) * q
	np.USD = np.USD + Money(aux)
}

// Operation type to specify selling of buying
type Operation string

func (o Operation) Equal(val string) bool {
	return strings.ToUpper(string(o)) == val
}

const (
	Buy  Operation = "buy"
	Sell Operation = "sell"
)

// Trd is a Trade
type Trd struct {
	ID     string
	Type   string
	USD    int64
	Gold   float64
	OurT   string
	TheirT string
}

// GetNPos operate over every trade and return a NetPos responding
// for every buying and selling operation
func GetNPos(trades []Trd) NetPos {
	var np NetPos

	for _, t := range trades {
		if Buy.Equal(t.Type) {
			np.AddGold(t.Gold, t.USD)
		}
		if Sell.Equal(t.Type) {
			np.SellGold(t.Gold, t.USD)
		}
	}

	return np
}

// Node represent a node in a data struct with a generic value
type Node struct {
	Val   Trd
	Nodes []*Node
}

// AddNode add neighborhoods to the node
func (n *Node) AddNode(node *Node) {
	if n.contains(node) {
		return
	}
	n.Nodes = append(n.Nodes, node)
}
func (n *Node) contains(node *Node) bool {
	for _, n := range n.Nodes {
		if n == node {
			return true
		}
	}
	return false
}

type Graph struct {
	Nodes []*Node
}

func (g *Graph) AddNode(val Trd) {
	node := &Node{Val: val}
	if len(g.Nodes) == 0 {
		g.Nodes = append(g.Nodes, node)
		return
	}

	for _, n := range g.Nodes {
		nde := n
		if nde.Val.OurT == val.OurT {
			nde.AddNode(node)
			node.AddNode(nde)
		}
		if nde.Val.TheirT == val.TheirT {
			nde.AddNode(node)
			node.AddNode(nde)
		}
	}
	g.Nodes = append(g.Nodes, node)
}

// GetNodeByTag returns a Node containing a specific tag to
// could run through the graph
func (g *Graph) GetNodeByTag(tag string) *Node {
	for _, n := range g.Nodes {
		if n.Val.OurT == tag {
			return n
		}
		if n.Val.TheirT == tag {
			return n
		}
	}
	return nil
}

// visits contains all nodes who were visited by an iteration
type visits []*Node

func (v *visits) AddVisitor(node *Node) {
	*v = append(*v, node)
}
func (v *visits) WasVisited(node *Node) bool {
	for _, n := range *v {
		if n == node {
			return true
		}
	}
	return false
}

// DFS is Depths First Search to iterate recursively through a graph
func DFS(node *Node) []Trd {
	var trades []Trd
	trades = dfs(node, visits{}, trades)

	return trades
}
func dfs(n *Node, visits visits, trades []Trd) []Trd {
	if visits.WasVisited(n) {
		return trades
	}

	trades = append(trades, n.Val)

	visits.AddVisitor(n)
	for _, node := range n.Nodes {
		trades = dfs(node, visits, trades)
	}
	return trades
}

// MainTrd receive trades and tags and create a graph
// Later call DFS for every tag to create a map
func MainTrd(trades []Trd, tag string) []Trd {
	// here we can create the graph with trades and tags and their connections
	// everything is a node

	graph := Graph{Nodes: make([]*Node, 0)}

	for _, trade := range trades {
		graph.AddNode(trade)
	}

	node := graph.GetNodeByTag(tag)

	return DFS(node)
}
