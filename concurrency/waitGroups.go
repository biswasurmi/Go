// ✅ What This Code Does
// It starts 5 goroutines (workers).

// Each worker sends a message to a channel.

// A separate goroutine reads from the channel and prints the messages.

// The WaitGroup ensures the main function waits for all workers to finish before closing the channel.

package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	msg := fmt.Sprintf("worker %d: task done", id)
	ch <- msg
}
func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	n := 5

	go func() {
		for msg := range ch {
			fmt.Println(msg)
		}
	}()
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, &wg, ch)
	}
	wg.Wait()
	close(ch)
}

// Main → Starts reader goroutine
//      → Starts 5 worker goroutines (one by one)
//      → Waits for all workers to call Done()
//      → Closes the channel

// Each Worker → Sends message to channel → Calls Done()

// Reader Goroutine → Reads & prints each message → Stops when channel closes
