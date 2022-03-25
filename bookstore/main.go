package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ajay-code/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	fmt.Println("server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}