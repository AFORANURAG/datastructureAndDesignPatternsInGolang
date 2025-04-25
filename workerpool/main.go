package main

// lets implement a worker pool
// we will have a pool of workers
import (
	"fmt"
	"time"
)

// id of the job
// jobs channel
// results channel
func Worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	fmt.Println("Hello, World!")
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 0; w < 10; w++ {
		go Worker(w, jobs, results)
	}
	for j := 1; j <= 100; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}
