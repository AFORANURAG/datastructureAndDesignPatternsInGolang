package implementations

import "fmt"

func BuildGraphWithAdjacencyMatrix() {
	vertices := 4
	graph := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		graph[i] = make([]int, vertices)
	}
	addEdgeMatrix(graph, 0, 1)
	addEdgeMatrix(graph, 0, 2)
	addEdgeMatrix(graph, 0, 3)
	addEdgeMatrix(graph, 2, 3)

	for _, row := range graph {
		fmt.Println(row)
	}

}

func addEdgeMatrix(graph [][]int, u, v int) {
	// when i say add a edge between u and v(bidirectionally) that means
	graph[u][v] = 1
	graph[v][u] = 1
}
