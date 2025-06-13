package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; i < 5; i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string) {
	for i := 0; i < 5; i++ {
		c <- "pong"
		// msg <- c will show error as send only
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		// c <- msg will show error as receive only
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}

/*
| Condition                                     | Why no deadlock?                   |
| --------------------------------------------- | ---------------------------------- |
| Unbuffered channel                            | Requires sync between send/receive |
| One receiver (`printer`) alive                | Always receiving from channel      |
| Senders (`pinger`, `ponger`) run concurrently | No one gets stuck waiting forever  |
| `main()` stays alive with `Scanln`            | Prevents premature program exit    |
*/
