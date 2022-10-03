package parser

import (
	"reflect"
	"strings"
)

// Money approach is used in order to not use float64, since this kind of operations
// generate bugs in floating points. So the base money is a integer
type Money int64

const (
	USD  Money = 100000000
	GOLD       = USD * 1750
)

// NetPos is the net position
type NetPos struct {
	Gold float32
	USD  Money
}

// AddGold add gold quantity and also the operation to subtract usd
func (np *NetPos) AddGold(q float32, total int64) {
	np.Gold += q

	aux := float32(GOLD) * q
	np.USD = np.USD - Money(aux)
}

// SellGold subtract gold from net position and also sum usd
func (np *NetPos) SellGold(q float32, total int64) {
	np.Gold -= q

	aux := float32(GOLD) * q
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
	Gold   float32
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
	Val   interface{}
	Nodes []*Node
}

// AddNode add neighborhoods to the node
func (n *Node) AddNode(node *Node) {
	n.Nodes = append(n.Nodes, node)
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

type Tags map[string]Trd

func (t Tags) AddTrd(tag string, trd Trd) {
	t[tag] = trd
}

// DFS is Depths First Search to iterate recursive through a graph
func DFS(tag *Node, m map[string][]Trd) {
	var trades []Trd
	trades = dfs(tag, visits{}, trades)

	m[tag.Val.(string)] = trades
}
func dfs(n *Node, visits visits, trades []Trd) []Trd {
	if visits.WasVisited(n) {
		return trades
	}

	if t := reflect.TypeOf(n.Val); t.Kind() == reflect.Struct {
		trades = append(trades, n.Val.(Trd))
	}

	visits.AddVisitor(n)
	for _, node := range n.Nodes {
		trades = dfs(node, visits, trades)
	}
	return trades
}

// MainTrd receive trades and tags and create a graph
// Later call DFS for every tag to create a map
func MainTrd(trades []Trd, tags []string) map[string][]Trd {
	// here we can create the graph with trades and tags and their connections
	// everything is a node

	// create nodes from trades
	var trds []*Node
	for _, trade := range trades {
		trds = append(trds, &Node{Val: trade, Nodes: make([]*Node, 0)})
	}

	// create nodes from tags
	var tgs []*Node
	for _, tag := range tags {
		tgs = append(tgs, &Node{Val: tag, Nodes: make([]*Node, 0)})
	}

	// connect trades with tags
	for _, trade := range trds {
		t := trade.Val.(Trd)
		for _, tag := range tgs {
			if t.OurT == tag.Val.(string) || t.TheirT == tag.Val.(string) {
				trade.AddNode(tag)
				tag.AddNode(trade)
			}
		}
	}

	// create map response
	//add for every tag key it's trades
	res := make(map[string][]Trd)
	for _, tag := range tgs {
		DFS(tag, res)
	}

	return res
}
