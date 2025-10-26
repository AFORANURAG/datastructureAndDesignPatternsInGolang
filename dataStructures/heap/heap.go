package main

import (
	"fmt"
)

// =====================
// STRUCT DEFINITION
// =====================

type Heap struct {
	arr       []int
	isMinHeap bool // if true => MinHeap, else MaxHeap
}

// =====================
// HELPER FUNCTIONS
// =====================

func (h *Heap) compare(a, b int) bool {
	if h.isMinHeap {
		return a < b
	}
	return a > b
}

// Restores heap property from a given index downwards
func (h *Heap) heapify(n, i int) {
	largestOrSmallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h.compare(h.arr[left], h.arr[largestOrSmallest]) {
		largestOrSmallest = left
	}
	if right < n && h.compare(h.arr[right], h.arr[largestOrSmallest]) {
		largestOrSmallest = right
	}

	if largestOrSmallest != i {
		h.arr[i], h.arr[largestOrSmallest] = h.arr[largestOrSmallest], h.arr[i]
		h.heapify(n, largestOrSmallest)
	}
}

func (h *Heap) buildHeap() {
	n := len(h.arr)
	for i := n/2 - 1; i >= 0; i-- {
		h.heapify(n, i)
	}
}

func (h *Heap) insert(key int) {
	h.arr = append(h.arr, key)
	i := len(h.arr) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if h.compare(h.arr[parent], h.arr[i]) {
			break
		}
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		i = parent
	}
}

func (h *Heap) deleteRoot() int {
	n := len(h.arr)
	if n == 0 {
		panic("heap is empty")
	}
	root := h.arr[0]
	h.arr[0] = h.arr[n-1]
	h.arr = h.arr[:n-1]
	h.heapify(len(h.arr), 0)
	return root
}

func (h *Heap) heapSort() []int {
	n := len(h.arr)
	h.buildHeap() // Ensure valid heap before sorting

	for i := n - 1; i > 0; i-- {
		h.arr[0], h.arr[i] = h.arr[i], h.arr[0]
		h.heapify(i, 0)
	}
	return h.arr
}

func main() {
	// Example 1: Min Heap
	fmt.Println("---- MIN HEAP ----")
	minHeap := Heap{arr: []int{4, 10, 3, 5, 1}, isMinHeap: true}
	minHeap.buildHeap()
	fmt.Println("After buildHeap:", minHeap.arr)

	minHeap.insert(2)
	fmt.Println("After insert(2):", minHeap.arr)

	fmt.Println("Deleted root:", minHeap.deleteRoot())
	fmt.Println("After deleteRoot:", minHeap.arr)

	// Example 2: Max Heap + HeapSort
	fmt.Println("\n---- MAX HEAP ----")
	maxHeap := Heap{arr: []int{4, 10, 3, 5, 1}, isMinHeap: false}
	maxHeap.buildHeap()
	fmt.Println("After buildHeap:", maxHeap.arr)

	sorted := maxHeap.heapSort()
	fmt.Println("After heapSort:", sorted)
}
