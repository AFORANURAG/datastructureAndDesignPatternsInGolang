package implementations

func BuildGraphWithAdjacencyList() map[int][]int {
	vertices := 4

	graph := make(map[int][]int, vertices)

	// key(node):[node connected to key node ]
	//   0
	// 1  2---3
	//    0 connected bidirectionally to 1 and 2
	// 2 connected bidirectionally to 3

	addEdgeList(graph, 0, 1)
	addEdgeList(graph, 1, 3)
	addEdgeList(graph, 2, 0)
	addEdgeList(graph, 2, 3)

	return graph
}

func addEdgeList(graph map[int][]int, u, v int) {
	graph[u] = append(graph[u], v)
	graph[v] = append(graph[v], u) // undirected
}
