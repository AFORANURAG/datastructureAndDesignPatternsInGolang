package implementations

import "fmt"

func Dfs(
	node int,
	graph map[int][]int,
	visited map[int]bool,
) {

	visited[node] = true

	fmt.Println(node)

	for _, nei := range graph[node] {

		if !visited[nei] {
			Dfs(nei, graph, visited)
		}
	}
}
