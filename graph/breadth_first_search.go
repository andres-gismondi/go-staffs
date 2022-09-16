package graph

import "fmt"

// BFS graph goes over every node, so you need to divide the vertices into two categories:
// Visited and Not Visited with an array of booleans
// BFS uses a queue data structure for traversal where enqueue the neighbours not visited
func BFS(n Node) {
	var vi visits
	q := Queue{}

	q.Add(n)

	for !q.Empty() {
		node := q.First()
		fmt.Println(node.ID)
		vi = append(vi, visited{NodeID: n.ID, Value: true})

		for _, nn := range node.Neighbours {
			if ok := vi.WasVisited(nn.ID); !ok {
				q.Add(nn)
			}
		}
	}
}
