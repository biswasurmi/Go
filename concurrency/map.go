package main

import (
	"fmt"
	"sync"
)

// sync.map => load, store, delete
func main() {
	var m sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key %d", id)
			value := fmt.Sprintf("value %d", id)
			m.Store(key, value)
		}(i)
		wg.Wait()

		for i := 0; i < 5; i++ {
			key := fmt.Sprintf("key %d", i)
			if val, ok := m.Load(key); ok {
				fmt.Printf("key : %s, value: %s \n", key, val)
			}
		}
	}
}
