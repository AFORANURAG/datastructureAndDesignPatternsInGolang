package implementations

func HasCycleDirectedDFS(node int, graph map[int][]int, visited, pathVisited map[int]bool) bool {
	visited[node] = true
	pathVisited[node] = true

	for _, nei := range graph[node] {
		if !visited[nei] {

			isCyclic := HasCycleDirectedDFS(nei, graph, visited, pathVisited)
			if isCyclic {
				return true
			}
		} else {
			// visited
			if pathVisited[nei] {
				return true
			}
		}
	}
	pathVisited[node] = false
	return false
}
