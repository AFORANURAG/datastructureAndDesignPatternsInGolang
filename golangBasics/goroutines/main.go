package main

import (
	"fmt"
	"time"
)

func func1(ch chan string) {
	ch <- "boom"
}

func main() {
	// channel is a way to communicate between goroutines
	// they are synchronization mechanism
	// they are used to send and receive data between goroutines
	// they are type safe
	// they are blocking by default
	// they are unbuffered by default
	// they are synchronous by default
	// they are unidirectional by default
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func1(ch1)
	go func1(ch2)
	count := 0
	for count < 3 {
		select {
		case msg1 := <-ch1:
			fmt.Println("ch1 received", msg1)
			break
		case msg2 := <-ch2:
			fmt.Println("ch2 received", msg2)
			break
		default:
			fmt.Println("no message received")
		}
		time.Sleep(time.Millisecond)
		count++
	}
}



package main

import (
	"fmt"
)

// Node structure
type Node struct {
	data  int
	left  *Node
	right *Node
}

