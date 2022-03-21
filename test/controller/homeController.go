package controller

import (
	"fmt"
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to API by ajay singh")
}
