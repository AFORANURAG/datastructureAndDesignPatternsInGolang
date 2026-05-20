package implementations

import "fmt"

func Bfs(start int, graph map[int][]int) {

	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println("node is", node)
		for _, nei := range graph[node] {
			if !visited[nei] {
				visited[nei] = true
				queue = append(queue, nei)
			}
		}
	}

}
