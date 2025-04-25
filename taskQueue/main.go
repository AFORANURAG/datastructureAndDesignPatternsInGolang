package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID     int
	Data   interface{}
	Result chan interface{}
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	taskQueue := make(chan Task, numWorkers)
	// channels are just like pipes, they connect the workers and the main function
	// they are queues, they are FIFO
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, taskQueue, &wg)
	}
	for i := 0; i < 10; i++ {
		resultChan := make(chan interface{})
		taskQueue <- Task{ID: i, Data: fmt.Sprintf("Task %d", i), Result: resultChan}
		go func(id int, resChan chan interface{}) {
			result := <-resChan
			fmt.Printf("Task %d result: %v\n", id, result)
		}(i, resultChan)

	}
	close(taskQueue)
	wg.Wait()
	fmt.Println("All tasks completed")
}
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)

		time.Sleep(100 * time.Millisecond)
		// Process task and send result
		result := fmt.Sprintf("Processed: %v", task.Data)
		task.Result <- result
	}
	fmt.Printf("Worker %d shutting down\n", id)
}
