package main

import (
	"fmt"
	"net/http"
	"test/route"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	route.AddRoutes(r)
	fmt.Println("Running server on http://localhost:5000")
	http.ListenAndServe(":5000", r)

}
