package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Request struct {
	data int
	resp chan float64
}

func CreateAndRequest(req chan Request) {
	resp := make(chan float64)
	// spawn requests indefinitely
	for {
		// wait before next request
		time.Sleep(time.Duration(rand.Int63n(int64(time.Millisecond))))
		req <- Request{int(rand.Int31n(90)), resp}
		// read value from RESP channel
		<-resp
	}
}

type Work struct {
	// heap index
	idx int

	// central worker channel from which requests are received
	wok chan Request

	pending int
}

func (w *Work) Serve(done chan *Work) {
	// so it is a worker that is serving requests
	// it is forever waiting for request via w.wok channel that is a central channel
	// when it receives a request it serves it and sends the response back via the resp channel in the Request struct
	for {
		req := <-w.wok
		req.resp <- math.Sin(float64(req.data))
		done <- w
	}
}

// worker pool
type Pool []*Work

type Balancer struct {
	pool Pool
	done chan *Work
}

func InitBalancer() *Balancer {
	done := make(chan *Work, 3)
	b := &Balancer{pool: make(Pool, 0, 3), done: done}
	for i := 0; i < 3; i++ {
		w := &Work{idx: i, wok: make(chan Request)}
		go w.Serve(done)
		heap.Push(&b.pool, w)
	}
	return b
}

func (b *Balancer) Balance(req chan Request) {

	for {
		select {
		case request := <-req:
			b.dispath(request)

		case w := <-b.done:
			b.completed(w)
		}
		b.print()
	}

}

// heap is used to get the worker with the least pending requests, that it.
// so we can dispatch the request to the worker with the least pending requests
// and we can also push the worker back to the heap when it is done
// so that it can be used to serve other requests

func (b *Balancer) dispath(req Request) {
	w := heap.Pop(&b.pool).(*Work)
	w.wok <- req
	w.pending++
	heap.Push(&b.pool, w)
}

// what would be completed function

func (b *Balancer) completed(w *Work) {
	w.pending--
	// remove from heap, why are we doing this?
	// because we want to remove the worker from the heap when it is done
	// so that it can be used to serve other requests
	// so we are not using the same worker to serve the same request again
	// and again

	heap.Remove(&b.pool, w.idx)

	heap.Push(&b.pool, w)
}

func (b *Balancer) print() {
	sum := 0
	sumsq := 0
	// Print pending stats for each worker
	for _, w := range b.pool {
		fmt.Printf("%d ", w.pending)
		sum += w.pending
		sumsq += w.pending * w.pending
	}
	// Print avg for worker pool
	avg := float64(sum) / float64(len(b.pool))
	variance := float64(sumsq)/float64(len(b.pool)) - avg*avg
	fmt.Printf(" %.2f %.2f\n", avg, variance)
}

func main() {
	work := make(chan Request)
	for i := 0; i < 3; i++ {
		go CreateAndRequest(work)
	}
	InitBalancer().Balance(work)
}

// heap.Interface implementation

func (p Pool) Len() int { return len(p) }

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p *Pool) Swap(i, j int) {
	a := *p
	a[i], a[j] = a[j], a[i]
	a[i].idx = i
	a[j].idx = j
}

func (p *Pool) Push(x interface{}) {
	n := len(*p)
	item := x.(*Work)
	item.idx = n
	*p = append(*p, item)
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	item.idx = -1 // for safety
	*p = old[0 : n-1]
	return item
}
