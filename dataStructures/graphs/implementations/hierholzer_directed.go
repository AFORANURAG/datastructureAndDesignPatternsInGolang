package implementations

import "fmt"

type Edge struct {
	To     int
	Weight int // ignored by Hierholzer
}

type Graph struct {
	Adj [][]Edge
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func (g *Graph) Hierholzer(start int) []int {
	path := []int{}

	var dfs func(int)

	dfs = func(u int) {

		for len(g.Adj[u]) > 0 {

			// Pop last edge
			last := len(g.Adj[u]) - 1
			edge := g.Adj[u][last]
			g.Adj[u] = g.Adj[u][:last]

			dfs(edge.To)
		}

		path = append(path, u)
	}

	dfs(start)

	reverse(path)

	return path
}

func main() {

	g := Graph{
		Adj: make([][]Edge, 4),
	}

	g.Adj[0] = append(g.Adj[0], Edge{To: 1})
	g.Adj[1] = append(g.Adj[1], Edge{To: 2})
	g.Adj[2] = append(g.Adj[2], Edge{To: 0})
	g.Adj[0] = append(g.Adj[0], Edge{To: 3})

	path := g.Hierholzer(0)

	fmt.Println(path)
}
