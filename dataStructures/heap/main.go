package main

import (
	"fmt"
	impl "heap/implementation"
)

func main() {
	h := &impl.MaxHeap{}
	h.Push(10)
	h.Push(4)
	h.Push(15)
	h.Push(20)
	h.Push(0)
	h.Delete(3)

	fmt.Println("Heap array:", h.Arr)

	fmt.Println("Pop:", h.Pop())
	fmt.Println("Pop:", h.Pop())
	fmt.Println("Heap array:", h.Arr)

	fmt.Println("\n\n<-----------------------------Heapified array:-------------------------------->\n\n")
	a := []int{5, 8, 5}
	a = impl.Heapify(a)
	fmt.Println("Heapified arr is", a)
	a = impl.HeapSort(a)
	fmt.Println("SortedArr is", a)

}
