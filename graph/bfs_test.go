package graph_test

import (
	"testing"

	"go-staffs/graph"
)

func TestGraph_BFS(t *testing.T) {
	n0 := graph.NewNode(0, 0)
	n1 := graph.NewNode(1, 1)
	n2 := graph.NewNode(2, 2)
	n3 := graph.NewNode(3, 3)

	n0.AddNeighbour(n1)
	n0.AddNeighbour(n2)

	n1.AddNeighbour(n2)

	n2.AddNeighbour(n0)
	n2.AddNeighbour(n3)

	n3.AddNeighbour(n3)
	n3.AddNeighbour(n2)

	graph.BFS(n2)
}
