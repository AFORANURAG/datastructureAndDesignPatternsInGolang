package heap

type MaxHeap struct {
	Arr []int
}

func Parent(i int) int { return ((i - 1) / 2) }
func Left(i int) int   { return (2*i + 1) }
func Right(i int) int  { return (2*i + 2) }

func (h *MaxHeap) Push(x int) {
	h.Arr = append(h.Arr, x)
	h.siftUp(len(h.Arr) - 1)
}
func (h *MaxHeap) siftUp(i int) {
	for i != 0 && h.Arr[Parent(i)] < h.Arr[i] {
		h.Arr[Parent(i)], h.Arr[i] = h.Arr[i], h.Arr[Parent(i)]
		i = Parent(i)
	}
}
func (h *MaxHeap) Pop() int {
	if len(h.Arr) == 0 {
		panic("heap is empty")
	}
	root := h.Arr[0]
	last := h.Arr[len(h.Arr)-1]
	h.Arr[0] = last
	h.Arr = h.Arr[:len(h.Arr)-1]
	h.siftDown(0)
	return root
}
func (h *MaxHeap) Update(i int, newVal int) {
	val := h.Arr[i]
	h.Arr[i] = newVal
	if newVal > val {
		h.siftUp(i)
	} else {
		h.siftDown(i)
	}

}

func (h *MaxHeap) Delete(i int) {
	n := len(h.Arr)
	if i >= n {
		return
	}
	h.Arr[i] = h.Arr[n-1]
	h.Arr = h.Arr[:n-1]
	h.siftDown(i)
	h.siftUp(i)
}

func (h *MaxHeap) siftDown(i int) {
	smallest := i
	right := Right(i)
	left := Left(i)

	if left < len(h.Arr) && h.Arr[left] < h.Arr[smallest] {
		smallest = left
	}
	if right < len(h.Arr) && h.Arr[right] < h.Arr[smallest] {
		smallest = right
	}
	if i != smallest {
		h.Arr[i], h.Arr[smallest] = h.Arr[smallest], h.Arr[i]
		h.siftDown(smallest)
	}
}

func (h *MaxHeap) Heapify() {
	n := len(h.Arr)
	for i := n/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
}

func siftDown(arr []int, i, size int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < size && arr[left] < arr[smallest] {
		smallest = left
	}
	if right < size && arr[right] < arr[smallest] {
		smallest = right
	}

	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		siftDown(arr, smallest, size)
	}
}

func Heapify(arr []int) []int {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(arr, i, n)
	}
	return arr
}

func HeapSort(arr []int) []int {
	n := len(arr)
	// heapify
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(arr, i, n)
	}

	for end := n - 1; end > 0; end-- {
		arr[0], arr[end] = arr[end], arr[0]
		siftDown(arr, 0, end)
	}
	return arr
}
