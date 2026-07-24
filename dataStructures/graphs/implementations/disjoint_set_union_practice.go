// lets implement a disjoint set union data structure

package implementations

type DisjointSetUnion struct {
	Parent []int
	Rank   []int
}

// initialize the DSU with n elements
func (dsu *DisjointSetUnion) NewDisjointSetUnion(n int) *DisjointSetUnion {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DisjointSetUnion{
		Parent: parent,
		Rank:   rank,
	}

}

func (dsu *DisjointSetUnion) Find(x int) int {
	// either the parent is the same as the element or some other element will be parent of the element
	// base case is when the parent is the same as the element
	// if not, we need to find the parent of the parent
	if dsu.Parent[x] != x {
		dsu.Parent[x] = dsu.Find(dsu.Parent[x])
	} // path compression
	return dsu.Parent[x]
}

// union by rank

func (dsu *DisjointSetUnion) Union(x, y int) bool {
	rootX := dsu.Find(x)
	rootY := dsu.Find(y)
	if rootX == rootY {
		return false
	}
	// let me write the algo

	if dsu.Rank[rootX] > dsu.Rank[rootY] {
		dsu.Parent[rootY] = rootX
	} else if dsu.Rank[rootX] < dsu.Rank[rootY] {
		dsu.Parent[rootX] = rootY
	} else {
		// if the ranks are the same, we can choose any one as the parent
		dsu.Parent[rootY] = rootX
		dsu.Rank[rootX]++
	}

	return true
}

func (dsu *DisjointSetUnion) IsConnected(x, y int) bool {
	return dsu.Find(x) == dsu.Find(y)
}
