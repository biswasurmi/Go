package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("serving at http://localhost:8081")
	err := http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("listen and serve tls error: ", err)
	}
}


// ✅ Step 1: Generate Self-Signed TLS Certificate
// Run the following in your terminal inside your project directory:

// bash
// Copy
// Edit
// openssl req -x509 -newkey rsa:2048 -keyout server.key -out server.crt -days 365 -nodes


// ✅ Step 3: Run the Server
// bash
// Copy
// Edit
// go run main.go
// ✅ Step 4: Open in Browser
// Go to:

// arduino
// Copy
// Edit
// https://localhost:8081/
// You’ll likely see a warning like “Your connection is not private” — because it’s self-signed.

// Click:

// Advanced → Proceed to localhost (unsafe)

// This is expected in local dev.