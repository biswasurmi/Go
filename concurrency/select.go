// | Use case               | Code pattern                 |
// | ---------------------- | ---------------------------- |
// | Wait on multiple chans | `select { case <-ch1, ... }` |
// | Add timeout            | `case <-time.After(...)`     |
// | Avoid blocking         | `default:`                   |
// | Random choice on ready | Multiple ready → random pick |

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- "from ch1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "from ch2"
	}()
	for i := 0; i < 2; i++ {
		select {
		// case msg1 := <- ch1:
		// 	fmt.Println("Received:", msg1)
		// case msg2 := <- ch2:
		// 	fmt.Println("Received:", msg2)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
		default:
			fmt.Println("Nothing is ready")
		}

	}
}
