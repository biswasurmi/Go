package main

import (
	"fmt"
	"sync"
)

// memory

// sync.pool => .GET , .PUT

type someObject struct {
	data []byte
}

func createObject() *someObject {
	return &someObject{
		data: make([]byte, 1024*1024),
	}
}

func main() {
	var memoryPiece int
	//var objects []*someObject
	objectPool := sync.Pool{
		New: func() interface{} {
			memoryPiece++
			return createObject()
		},
	}
	const worker = 1024 * 1024
	var wg sync.WaitGroup
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			obj := objectPool.Get().(*someObject)
			objectPool.Put(obj)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Done", memoryPiece)
}
