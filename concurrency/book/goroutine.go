package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 3; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(10))
		time.Sleep(amt * time.Millisecond)
	}
}

func main() {
	for i := 0; i < 3; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input) // scanln is used to prevent main to close after it finishes its work, through scanln the program waits for an enter press...
}
