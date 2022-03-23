package main

import (
	"fmt"
	"net/http"
	"server/pkg/router"
)

func main() {
	r := router.Router()

	fmt.Println("server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}