package implementations

// visited and not parent

type Pair struct {
	node   int
	parent int
}

func HasCyclicUndirectedDFS(node, parent int, graph map[int][]int, visited map[int]bool) bool {
	visited[node] = true
	for _, nei := range graph[node] {
		if !visited[nei] {
			// call the HasCyclicUndirectedDFS
			isCyclic := HasCyclicUndirectedDFS(nei, node, graph, visited)
			if isCyclic {
				return true
			}
		} else {
			// visited
			if nei != parent {
				// in which case, it has a cycle
				return true
			}
		}

	}
	return false
}

func HasCyclicUndirectedBFS(start int, graph map[int][]int, visited map[int]bool) bool {
	queue := []Pair{{start, -1}}
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, nei := range graph[curr.node] {
			if !visited[nei] {
				// not visited
				visited[nei] = true
				queue = append(queue, Pair{nei, curr.node})
			} else {
				if curr.parent != nei {
					return true
				}

			}
		}
	}
	return false
}
