package graph_test

import (
	"testing"

	graph2 "go-staffs/data_structures/graph"
)

func TestGraph_DFS(t *testing.T) {
	n0 := graph2.NewNode(0, 0)
	n1 := graph2.NewNode(1, 1)
	n2 := graph2.NewNode(2, 2)
	n3 := graph2.NewNode(3, 3)

	n0.AddNeighbour(n1)
	n0.AddNeighbour(n2)

	n1.AddNeighbour(n2)

	n2.AddNeighbour(n0)
	n2.AddNeighbour(n3)

	n3.AddNeighbour(n3)
	n3.AddNeighbour(n2)

	graph2.DFS(n2)
}
