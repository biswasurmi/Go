package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		if r.URL.Path == "/" {
			path = "index.html"
		}
		http.ServeFile(w, r, path)
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Println("serving on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}



// ✅ 4. Run the Server
// In the terminal:

// bash
// Copy
// Edit
// go run main.go
// Output:

// nginx
// Copy
// Edit
// Serving on http://localhost:8081
// ✅ 5. Test in Browser
// Go to http://localhost:8081/index.html
// → Should render Hello World!

// Go to http://localhost:8081/edit.html
// → Should render content from edit.html

// Go to http://localhost:8081/hi
// → Responds with Hi

// Go to http://localhost:8081/
// → Also renders index.html by default