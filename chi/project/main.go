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

// Book represents the data model for a book.
type Book struct {
	UUID        string   `json:"uuid"`
	Name        string   `json:"name"`
	AuthorList  []string `json:"authorList"`
	PublishDate string   `json:"publishDate"`
	ISBN        string   `json:"isbn"`
}

var (
	tokenAuth  *jwtauth.JWTAuth             // JWT authentication handler
	adminUser  = "AdminUser"                // Basic auth username
	adminPass  = "AdminPassword"            // Basic auth password
	bookStore  = make(map[string]Book)      // In-memory book storage: map[UUID]Book
	authEnable bool                         // Flag to enable or disable authentication
)

func main() {
	// Initialize Chi router and add middleware for logging
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Endpoint to get JWT token using Basic Auth credentials
	r.Get("/api/v1/get-token", getTokenHandler)

	// Group routes under /api/v1/books
	r.Route("/api/v1/books", func(r chi.Router) {
		// If authentication enabled, require Basic Auth for listing books
		r.Group(func(r chi.Router) {
			if authEnable {
				r.Use(basicAuthMiddleware(adminUser, adminPass))
			}
			r.Get("/", listBooks)
		})

		// Routes that require JWT authentication
		r.Group(func(r chi.Router) {
			if authEnable {
				r.Use(jwtauth.Verifier(tokenAuth))      // Verify JWT token
				r.Use(jwtauth.Authenticator(tokenAuth)) // Reject unauthorized requests
			}
			r.Post("/", createBook)       // Create new book
			r.Get("/", listBooks)         // List all books (also accessible with JWT)
			r.Get("/{id}", getBook)       // Get book by UUID
			r.Put("/{id}", updateBook)    // Update book by UUID
			r.Delete("/{id}", deleteBook) // Delete book by UUID
		})
	})

	// Command line flags: auth enable/disable and port selection
	var port string
	flag.BoolVar(&authEnable, "auth", true, "Enable authentication")
	flag.StringVar(&port, "port", "8080", "Port to run the book server")
	flag.Parse()

	if !authEnable {
		fmt.Println("Authentication is disabled")
	}
	addr := fmt.Sprintf(":%s", port)
	fmt.Println("Starting server on port", port)

	// Start HTTP server
	http.ListenAndServe(addr, r)
}

// getTokenHandler provides a JWT token if valid Basic Auth credentials are provided
func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != adminUser || pass != adminPass {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	exp := time.Now().Add(100000 * time.Minute).Unix() // Token expiration time

	// Create JWT token with claims
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"user_id":  123,
		"username": adminUser,
		"exp":      exp,
	})
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Send token as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

// Initialize JWT auth with HS256 signing method and secret key
func init() {
	tokenAuth = jwtauth.New("HS256", []byte("supersecretkey123"), nil)
}

// basicAuthMiddleware protects endpoints with Basic Auth
func basicAuthMiddleware(username, password string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok || user != username || pass != password {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// createBook handles POST /api/v1/books to add a new book
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

// getBook handles GET /api/v1/books/{id} to fetch a book by UUID
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

// listBooks handles GET /api/v1/books to list all books
func listBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	for _, book := range bookStore {
		books = append(books, book)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// updateBook handles PUT /api/v1/books/{id} to update a book by UUID
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

// deleteBook handles DELETE /api/v1/books/{id} to remove a book by UUID
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
