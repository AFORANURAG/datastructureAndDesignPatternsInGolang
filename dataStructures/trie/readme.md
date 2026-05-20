//To use the heap package, you define a type that implements the following interface:

// type Interface interface {

//  sort.Interface          // Len(), Less(i, j int), Swap(i, j int)

//  Push(x any)             // add element x as Len() element

//  Pop() any               // remove and return last element

// }

import (

"container/heap"

)

typeIntHeap []int

func (h IntHeap) Len() int           { returnlen(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {

    *h = append(*h, x.(int))

}

func (h *IntHeap) Pop() any {

old := *h

n := len(old)

x := old[n-1]

    *h = old[0 : n-1]

return x

}
