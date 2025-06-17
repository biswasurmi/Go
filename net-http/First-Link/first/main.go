package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}



// ✅ Step 2: Initialize a Go Module (if not already done)
// In the same directory:

// bash
// Copy
// Edit
// go mod init mywebserver
// (You can name it anything.)

// ✅ Step 3: Run the Program
// bash
// Copy
// Edit
// go run main.go
// If it runs successfully, you’ll see no output, and the terminal will wait for requests.

// ✅ Step 4: Test in Browser or Curl
// Open your browser and go to:

// http://localhost:8081/hi → You'll see: Hi

// http://localhost:8081/hello/world → You'll see: Hello, "/hello/world"

// Or test via terminal:

// bash
// Copy
// Edit
// curl http://localhost:8081/hi
// curl http://localhost:8081/test
