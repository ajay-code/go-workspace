package main

import (
	"fmt"

	"mongo/router"
	"net/http"
)

func main() {
	port := ":5000"
	r := router.Router()

	fmt.Printf("server running on http://localhost%s", port)
	http.ListenAndServe(port, r)
}
