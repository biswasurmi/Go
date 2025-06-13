package main

import (
	"fmt"
	"sync"
)

type safecounter struct {
	mu  sync.Mutex
	val int
}

func (c *safecounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *safecounter) get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

func main() {
	var counter safecounter
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("final counter:", counter.get())
}
