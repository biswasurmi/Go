package main

import "fmt"

func main() {
	// chan := make(chan<- string) // only send
	// chan := make(<-chan string) // only receive
	// chan := make(chan string) // both send and receive

	ch1 := make(chan string)
	ch2 := make(chan string)

	// the below 2 lines mainly works parallely, before going to second line main gorouting knows that there is already a sending operation thus receiving won't block anyone and thus no deadlock
	// but if that's not the case, then second line will block main goroutine as it has no info about sending operation that's why this order should be maintained
	go sending(ch1)
	value := <-ch1
	fmt.Println("valueFromChannel1", value)
	//same as previous, order is important as main knows there is both sending and receiving option
	go receiving(ch2)
	ch2 <- value

	/******* unidirectional **************/
	channel1 := make(chan string)
	go convert(channel1)
	fmt.Println(<-channel1)
}
func convert(s chan<- string) {
	s <- "go go"
}
func sending(s chan string) {
	s <- "go ch1"
}
func receiving(s chan string) {
	fmt.Println(<-s)
}
