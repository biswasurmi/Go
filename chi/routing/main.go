package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// Named param: /users/{userID}
	r.Get("/users/{userID}", getUser)

	// Multiple named params with separator: /articles/{date}-{slug}
	r.Get("/articles/{date}-{slug}", getArticle)

	// Wildcard param: /static/*
	r.Get("/static/*", serveStatic)

	// Start server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

// /users/{userID}
func getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Write([]byte(fmt.Sprintf("User ID: %s\n", userID)))
}

// /articles/{date}-{slug}
func getArticle(w http.ResponseWriter, r *http.Request) {
	date := chi.URLParam(r, "date")
	slug := chi.URLParam(r, "slug")
	w.Write([]byte(fmt.Sprintf("Article Date: %s\nSlug: %s\n", date, slug)))
}

// /static/*
func serveStatic(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "*")
	w.Write([]byte(fmt.Sprintf("Serving static file: %s\n", path)))
}
