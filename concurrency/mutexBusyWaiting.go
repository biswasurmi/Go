package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.Mutex
	condVar *sync.Cond
	data    int
)

func main() {
	condVar = sync.NewCond(&mu)
	go producer()
	go consumer()
	time.Sleep(5 * time.Second)
}

func producer() {
	for i := 0; i < 5; i++ {
		mu.Lock()
		data = i
		condVar.Signal()
		mu.Unlock()
		time.Sleep(time.Second)
	}
}

func consumer() {
	counter := 0
	for i := 0; i < 5; i++ {
		mu.Lock()
		for data != i {
			counter++
			condVar.Wait()
		}
		fmt.Println(data, counter)
		mu.Unlock()
	}
}
