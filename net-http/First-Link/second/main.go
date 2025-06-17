package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)


var counter int

var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request){
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()	
}

func main() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}


// 🚀 How to Run Step-by-Step
// ✅ Step 1: Save Code to File
// Save the code to a file named main.go.

// ✅ Step 2: Initialize Module (if not done)
// bash
// Copy
// Edit
// go mod init counter-server
// ✅ Step 3: Run the Program
// bash
// Copy
// Edit
// go run main.go
// ✅ Step 4: Test Routes
// In browser or with curl:

// URL	Response	Description
// http://localhost:8081/	hello	Root route
// http://localhost:8081/hi	Hi	Simple greeting
// http://localhost:8081/increment	1, 2, 3, ...	Increments counter each visit