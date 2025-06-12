package main

import "fmt"

func main() {
	// channel type
	// var channel_name chan Type
	// channel_name := make(chan Type)

	//var firstchan chan int
	// firstchan := make(chan int)
	// fmt.Println("value of the channel:", firstchan)
	// fmt.Printf("Type of the channel: %T", firstchan)

	// *****************************channel creation, assigning******************************
	// chan1 <- value //send
	// var := <- chan1 // receive
	// <- chan1 // print

	// fmt.Println("Hello from main")
	// ch := make(chan int)
	// // ch <- 10 if we put this line before calling goroutine
	// // ❓ What Happens Here?
	// // ch := make(chan int) — a synchronous, unbuffered channel is created.

	// // ch <- 10 — this sends 10 into the channel, but blocks until another goroutine receives it.

	// // But at this point, there is no goroutine running yet to receive that value!

	// // So the program blocks indefinitely at ch <- 10 → Deadlock.

	// // go multiplyWithChan(ch) is never even reached until the current line completes (but it doesn't, due to the blocking send).
	// go multiplyWithChan(ch)
	// ch <- 10
	// // 🔁 Execution Flow (Working):
	// // ch := make(chan int) — creates the channel.

	// // go multiplyWithChan(ch) — spawns a goroutine that starts waiting to receive from the channel.

	// // ch <- 10 — main goroutine sends 10, and it's received immediately by the goroutine.

	// // The goroutine receives 10, multiplies it by 1000, and prints the result.

	// // main() continues and prints "Bye from main".
	// fmt.Println("Bye from main")

	// // ***************close and ok syntax**********************
	// ch := make(chan int)
	// // elem has the channel value and ok = false means channel is closed or vice versa
	// elem, ok := <- ch // this opens the channel but as the channel doesn't have any value it blocks here and creates deadlock
	// close(ch)
	// fmt.Println("Hello Go", ok, elem)

	// // close and ok syntax
	// ch := make(chan int)
	// close(ch) // close the channel
	// elem, ok := <- ch // immediately returns elem = 0 and ok = false
	// fmt.Println("Hello Go", ok, elem)

	/*
		| Channel State                       | Result of `<-ch`    | `ok` Value | Explanation                         |
		| ----------------------------------- | ------------------- | ---------- | ----------------------------------- |
		| Channel open, has value             | gets value          | `true`     | Normal receive                      |
		| Channel open, no value              | blocks              | -          | Waits until sender sends            |
		| Channel closed, values already sent | gets sent value     | `true`     | Drains sent values even after close |
		| Channel closed, no values left      | zero value (e.g. 0) | `false`    | Receive from closed, empty channel  |
	*/

	// ************* loop ****************

	/*
		🔑 Core Principle:
		If your goroutines communicate using channels (or other synchronization tools), you don't need time.Sleep() to wait for them.

		But...

		If there is no communication or coordination mechanism (like channels, WaitGroup, etc.), then time.Sleep() is often misused as a hack to "keep main alive".
	*/
	// ch := make(chan string)
	// go initString(ch)
	// for{
	// 	response, ok := <- ch
	// 	if ok == false {
	// 		fmt.Println("channel close", ok)
	// 		break
	// 	}
	// 	fmt.Println("channel open", response, ok)
	// }

	//************************** find length of channels *********************************
	// buffered channel => channels which have length
	// unbuffered channel => channels which doesn't have length and must have a receiver

	// ch := make(chan string, 5)
	// ch <- "abc"
	// ch <- "cde"
	// ch <- "efg"
	// ch <- "ghi"

	// fmt.Println("Length of the channel is:", len(ch))
	// fmt.Println("Capacity of the channel is:", cap(ch))

	/***************** print channel **************************/
	testch := make(chan string)
	go func() {
		testch <- "value 1"
		testch <- "value 2"
		testch <- "value 3"
		close(testch)
	}()
	for value := range testch {
		fmt.Println(value)
	}
}
func initString(ch chan string) {
	for i := 0; i < 3; i++ {
		ch <- "go go"
	}
	close(ch)
}

func multiplyWithChan(ch chan int) {
	fmt.Println(<-ch * 1000)
}
