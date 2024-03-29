package graph

import "fmt"

/*
Depth-first search is an algorithm for traversing or searching tree or graph data structures.
The algorithm starts at the root node (selecting some arbitrary node as the root node in the case of a graph)
and explores as far as possible along each branch before backtracking.

So the basic idea is to start from the root or any arbitrary node and mark the node and move
to the adjacent unmarked node and continue this loop until there is no unmarked adjacent node.
Then backtrack and check for other unmarked nodes and traverse them. Finally, print the nodes in the path.
*/

func DFS(node Node) {
	visits := visits{}
	dfs(node, visits)
}

func dfs(node Node, visits visits) {
	if ok := visits.WasVisited(node.ID); ok {
		return
	}

	fmt.Println(node.ID)
	visits = append(visits, visited{NodeID: node.ID, Value: true})
	for _, n := range node.Neighbours {
		dfs(n, visits)
	}
}
