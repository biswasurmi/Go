// Imagine multiple goroutines need to initialize a resource (e.g., load config, open a file, connect to DB). To make sure:

// That initialization happens only once

// It's thread-safe (no race conditions)

package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initialization done!")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(initialize)
	//initialize()
	fmt.Printf("worker %d is running\n", id)
}

func main() {
	var wg sync.WaitGroup
	n := 5

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
