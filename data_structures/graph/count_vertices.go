package graph

// CountVertices count the total number of ways or paths that exist between two vertices in a directed graph.
// These paths don’t contain a cycle, the simple enough reason is that a cycle contains an infinite number of paths
// and hence they create a problem
func CountVertices(init, dest Node) int {
	return count(init, dest, 0)
}

func count(actNode, desNode Node, counter int) int {
	if actNode.ID == desNode.ID {
		counter++
	}

	for _, n := range actNode.Neighbours {
		counter = count(n, desNode, counter)
	}

	return counter
}
