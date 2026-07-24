package main

import "fmt"

type Edge struct {
	To     int
	Weight int // ignored
	Id     int
}

type Graph struct {
	Adj      [][]Edge
	UsedEdge []bool
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func (g *Graph) AddEdge(u, v, w, id int) {

	g.Adj[u] = append(g.Adj[u], Edge{
		To:     v,
		Weight: w,
		Id:     id,
	})

	g.Adj[v] = append(g.Adj[v], Edge{
		To:     u,
		Weight: w,
		Id:     id,
	})
}

func (g *Graph) Hierholzer(start int) []int {

	path := []int{}

	var dfs func(int)

	dfs = func(u int) {

		for len(g.Adj[u]) > 0 {

			last := len(g.Adj[u]) - 1
			edge := g.Adj[u][last]

			// Pop
			g.Adj[u] = g.Adj[u][:last]

			if g.UsedEdge[edge.Id] {
				continue
			}

			g.UsedEdge[edge.Id] = true

			dfs(edge.To)
		}

		path = append(path, u)
	}

	dfs(start)

	reverse(path)

	return path
}

func main() {

	edgeCount := 5

	g := Graph{
		Adj:      make([][]Edge, 4),
		UsedEdge: make([]bool, edgeCount),
	}

	g.AddEdge(0, 1, 5, 0)
	g.AddEdge(1, 2, 8, 1)
	g.AddEdge(2, 0, 2, 2)
	g.AddEdge(0, 3, 4, 3)
	g.AddEdge(3, 2, 6, 4)

	path := g.Hierholzer(0)

	fmt.Println(path)
}
