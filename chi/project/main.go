package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

// Book model
type Book struct {
	UUID        string   `json:"uuid"`
	Name        string   `json:"name"`
	AuthorList  []string `json:"authorList"`
	PublishDate string   `json:"publishDate"`
	ISBN        string   `json:"isbn"`
}

var (
	tokenAuth *jwtauth.JWTAuth
	adminUser = "AdminUser"
	adminPass = "AdminPassword"
	bookStore = make(map[string]Book)
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	var authEnable bool
	r.Get("/api/v1/get-token", getTokenHandler)
	// Mount subrouter under /api/v1/books
	r.Route("/api/v1/books", func(r chi.Router) {
		//r.Use(basicAuthMiddleware(adminUser, adminPass)) // curl -u AdminUser:AdminPassword http://localhost:8080/api/v1/books for basic authentication

		if authEnable {
			r.Use(jwtauth.Verifier(tokenAuth))      // Verifies JWT from header/cookie
			r.Use(jwtauth.Authenticator(tokenAuth)) // Rejects unauthorized
		}

		r.Post("/", createBook)       // POST /api/v1/books
		r.Get("/", listBooks)         // GET /api/v1/books
		r.Get("/{id}", getBook)       // GET /api/v1/books/{id}
		r.Put("/{id}", updateBook)    // PUT /api/v1/books/{id}
		r.Delete("/{id}", deleteBook) // DELETE /api/v1/books/{id}
	})

	var port string

	flag.BoolVar(&authEnable, "auth", true, "Enable authentication")
	flag.StringVar(&port, "port", "8080", "Port to run the book server")
	flag.Parse()
	if !authEnable {
		fmt.Println("Authentication is disabled")
	}
	addr := fmt.Sprintf(":%s", port)
	fmt.Println("Starting server or port", port)

	http.ListenAndServe(addr, r)
}

// jwt token
func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != adminUser || pass != adminPass {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//exp := jwtauth.ExpireIn(100000 * 60) // returns int64
	exp := time.Now().Add(100000 * time.Minute).Unix()

	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"user_id":  123,
		"username": adminUser,
		"exp":      exp,
	})
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("supersecretkey123"), nil)
}

// func basicAuthMiddleware(username, password string) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			user, pass, ok := r.BasicAuth()
// 			if !ok || user != username || pass != password {
// 				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
// 				http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 				return
// 			}
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

// Create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusBadRequest)
		return
	}
	book.UUID = uuid.NewString()
	bookStore[book.UUID] = book

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Get a book by ID
func getBook(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "id")
	book, found := bookStore[bookID]
	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// List all books
func listBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	for _, book := range bookStore {
		books = append(books, book)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Update a book by ID
func updateBook(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "id")
	_, found := bookStore[bookID]
	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	var updatedBook Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request: %v", err), http.StatusBadRequest)
		return
	}

	updatedBook.UUID = bookID
	bookStore[bookID] = updatedBook

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// Delete a book by ID
func deleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "id")
	book, found := bookStore[bookID]
	if !found {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(bookStore, bookID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

/*

📚 Book REST API Documentation
✅ Overview
This API allows basic CRUD (Create, Read, Update, Delete) operations for books stored in memory (map[string]Book).

Each book has the following structure:

json
Copy
Edit
{
  "uuid": "auto-generated",
  "name": "string",
  "authorList": ["string"],
  "publishDate": "YYYY-MM-DD",
  "isbn": "string"
}
🚀 How to Run the Server
Make sure Go is installed.

Save your code to a file named main.go.

Run this in terminal:

bash
Copy
Edit
go run main.go
The server will start at:

arduino
Copy
Edit
http://localhost:8080
📘 Endpoints Summary
Method	Path	Description
POST	/api/v1/books	Create a new book
GET	/api/v1/books/{id}	Get a book by its UUID
GET	/api/v1/books	Get list of all books
PUT	/api/v1/books/{id}	Update a book by UUID
DELETE	/api/v1/books/{id}	Delete a book by UUID

📥 POST /api/v1/books
➤ Create a new book
Request:

bash
Copy
Edit
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Go Programming",
    "authorList": ["Alan A. A."],
    "publishDate": "2023-01-01",
    "isbn": "123-4567890123"
  }'
Response:

json
Copy
Edit
{
  "uuid": "generated-uuid",
  "name": "Go Programming",
  "authorList": ["Alan A. A."],
  "publishDate": "2023-01-01",
  "isbn": "123-4567890123"
}
📖 GET /api/v1/books/{id}
➤ Get a single book by its UUID
Example:

bash
Copy
Edit
curl http://localhost:8080/api/v1/books/your-book-uuid
Response:

json
Copy
Edit
{
  "uuid": "your-book-uuid",
  "name": "Go Programming",
  "authorList": ["Alan A. A."],
  "publishDate": "2023-01-01",
  "isbn": "123-4567890123"
}
If not found:

text
Copy
Edit
Book not Found
📚 GET /api/v1/books
➤ List all books
bash
Copy
Edit
curl http://localhost:8080/api/v1/books
Response:

json
Copy
Edit
[
  {
    "uuid": "uuid-1",
    "name": "Book 1",
    "authorList": ["Author A"],
    "publishDate": "2024-01-01",
    "isbn": "111-1111111111"
  },
  {
    "uuid": "uuid-2",
    "name": "Book 2",
    "authorList": ["Author B"],
    "publishDate": "2024-02-02",
    "isbn": "222-2222222222"
  }
]
🔁 PUT /api/v1/books/{id}
➤ Update an existing book
Example:

bash
Copy
Edit
curl -X PUT http://localhost:8080/api/v1/books/your-book-uuid \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Book",
    "authorList": ["New Author"],
    "publishDate": "2025-06-01",
    "isbn": "999-9999999999"
  }'
Response:

json
Copy
Edit
{
  "uuid": "your-book-uuid",
  "name": "Updated Book",
  "authorList": ["New Author"],
  "publishDate": "2025-06-01",
  "isbn": "999-9999999999"
}
❌ DELETE /api/v1/books/{id}
➤ Delete a book by its UUID
bash
Copy
Edit
curl -X DELETE http://localhost:8080/api/v1/books/your-book-uuid
Response:

json
Copy
Edit
{
  "uuid": "your-book-uuid",
  "name": "Deleted Book",
  "authorList": ["Author A"],
  "publishDate": "2024-01-01",
  "isbn": "111-1111111111"
}
If not found:

text
Copy
Edit
Book not found


w.Header().Set("Content-Type", "application/json")
🔍 What it does:
This tells the client (like curl, a browser, or Postman):

“The response body is in JSON format.”

It sets the HTTP response header:

pgsql
Copy
Edit
Content-Type: application/json
🤖 Why it's important:
Many clients use this header to decide how to parse the response.

Without this, the client might treat the response as plain text or HTML.

✅ This line is always recommended when you're sending JSON from a Go server.


json.NewEncoder(w).Encode(updatedBook)
🔍 What it does:
It creates a JSON encoder that writes directly to the http.ResponseWriter (w).

It serializes (marshals) updatedBook into JSON.

Then it writes that JSON into the response body.

This is equivalent to:

go
Copy
Edit
jsonBytes, _ := json.Marshal(updatedBook)
w.Write(jsonBytes)
But using json.NewEncoder(w).Encode(...) is:

Cleaner

More efficient (direct streaming)

Automatically writes a newline after the JSON


*/
