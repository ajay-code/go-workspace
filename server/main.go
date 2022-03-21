package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Book struct {
	Uuid uint
	Name string
}

func main() {

	books := []Book{{Uuid: 1, Name: "The legend of avatar"}, {Uuid: 2, Name: "Harry Potter"}}

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}).Methods("GET")

	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "api server")
	}).Methods("GET")

	r.HandleFunc("/api/books", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}).Methods("GET")

	r.HandleFunc("/api/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Fprintf(w, "book id: %s", id)

	}).Methods("GET")

	http.ListenAndServe("localhost:2000", removeTrailingSlash(r))
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
