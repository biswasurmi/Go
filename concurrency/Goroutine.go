package main

import (
	"fmt"
	"time"
)

func doSomeMagic() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hi", i)
	}
}
func main() { // main itself is a goroutine
	// go doSomeMagic() // this is a goroutine
	// doSomeMagic()
	//time.Sleep(1* time.Second) // if we don't execute this line then controller won't wait to execute goroutines

	/// anonymous go routine
	fmt.Println("Hello from main")

	func() {
		fmt.Println("Hello from anonymous function")
	}()
	go func() {
		fmt.Println("hello from anonymous goroutine")
	}()
	time.Sleep(1 * time.Second)
}
