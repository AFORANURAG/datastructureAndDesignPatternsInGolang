package implementations

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DSU{
		parent: parent,
		rank:   rank,
	}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
	x = d.Find(x)
	y = d.Find(y)
	if x == y {
		return false
	}
	if d.rank[x] < d.rank[y] {
		d.parent[x] = y
	} else if d.rank[x] > d.rank[y] {
		d.parent[y] = x
	} else {
		d.parent[y] = x
		d.rank[x]++
	}
	return true
}

func (d *DSU) IsConnected(x, y int) bool {
	return d.Find(x) == d.Find(y)
}
