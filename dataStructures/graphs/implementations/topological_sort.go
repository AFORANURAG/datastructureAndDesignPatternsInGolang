package implementations

func TopologicalSortDFS(graph map[int][]int) ([]int, bool) {
	nodes := allDirectedGraphNodes(graph)
	visited := make(map[int]bool)
	pathVisited := make(map[int]bool)
	order := make([]int, 0, len(nodes))

	var dfs func(node int) bool
	dfs = func(node int) bool {
		visited[node] = true
		pathVisited[node] = true

		for _, nei := range graph[node] {
			if !visited[nei] {
				if !dfs(nei) {
					return false
				}
			} else if pathVisited[nei] {
				return false
			}
		}

		pathVisited[node] = false
		order = append(order, node)
		return true
	}

	for node := range nodes {
		if !visited[node] {
			if !dfs(node) {
				return nil, false
			}
		}
	}

	reverseIntSlice(order)
	return order, true
}

func TopologicalSortBFS(graph map[int][]int) ([]int, bool) {
	nodes := allDirectedGraphNodes(graph)
	inDegree := make(map[int]int, len(nodes))

	for node := range nodes {
		inDegree[node] = 0
	}

	for _, neighbors := range graph {
		for _, nei := range neighbors {
			inDegree[nei]++
		}
	}

	queue := make([]int, 0, len(nodes))
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	order := make([]int, 0, len(nodes))
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)

		for _, nei := range graph[node] {
			inDegree[nei]--
			if inDegree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	if len(order) != len(nodes) {
		return nil, false
	}

	return order, true
}

func allDirectedGraphNodes(graph map[int][]int) map[int]struct{} {
	nodes := make(map[int]struct{})
	for node, neighbors := range graph {
		nodes[node] = struct{}{}
		for _, nei := range neighbors {
			nodes[nei] = struct{}{}
		}
	}
	return nodes
}

func reverseIntSlice(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
