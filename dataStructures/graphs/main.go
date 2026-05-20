package main

import (
	"fmt"
	impl "graphs/implementations"
)

func main() {
	fmt.Println("Hello graphs")
	graph := impl.BuildGraphWithAdjacencyList()
	visited := make(map[int]bool)
	for i := 0; i < len(graph); i++ {
		fmt.Printf("%d -> %v\n", i, graph[i])
	}
	impl.Bfs(0, graph)
	impl.Dfs(0, graph, visited)

	// impl.BuildGraphWithAdjacencyMatrix()
}
