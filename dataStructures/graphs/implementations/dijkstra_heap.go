package implementations

import (
	"container/heap"
	"fmt"
	"math"
)

// -------------------- Graph --------------------

type Edge struct {
	Vertex int
	Weight int
}

// -------------------- Priority Queue --------------------

type PriorityQueueItem struct {
	Vertex   int
	Distance int
}

type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance // Min Heap
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*PriorityQueueItem))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)

	item := old[n-1]
	old[n-1] = nil // avoid memory leak

	*pq = old[:n-1]

	return item
}

// -------------------- Dijkstra --------------------

func Dijkstra(n int, edges [][]int, src int) []int {

	// Build adjacency list
	adj := make([][]Edge, n)

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		adj[u] = append(adj[u], Edge{
			Vertex: v,
			Weight: w,
		})
	}
	// initialize distance array
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[src] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &PriorityQueueItem{
		Vertex:   src,
		Distance: 0,
	})

	for pq.Len() > 0 {
		// lets process
		item := heap.Pop(pq).(*PriorityQueueItem)
		u := item.Vertex
		d := item.Distance
		if d > dist[u] {
			continue
		}
		for _, edge := range adj[u] {
			v := edge.Vertex
			w := edge.Weight
			// update the estimates
			newDist := d + w
			if newDist < dist[v] {
				dist[v] = newDist
				heap.Push(pq, &PriorityQueueItem{Vertex: v, Distance: newDist})
			}

		}
	}
	return dist
}

// -------------------- Main --------------------

func main() {

	/*
		        0
		      /   \
		     4     1
		    /       \
		   1---------2
		      2

		Shortest from 0:
		0 -> 0 = 0
		0 -> 2 = 1
		0 -> 1 = 3
	*/

	edges := [][]int{
		{0, 1, 4},
		{0, 2, 1},
		{2, 1, 2},
	}

	n := 3

	ans := Dijkstra(n, edges, 0)

	fmt.Println(ans)
}
