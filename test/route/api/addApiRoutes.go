package api

import (
	"test/controller/course"

	"github.com/gorilla/mux"
)

func AddApiRoutes(r *mux.Router) {
	r.HandleFunc("/courses", course.GetAll).Methods("GET")
	r.HandleFunc("/courses", course.Create).Methods("POST")
	r.HandleFunc("/courses/{id}", course.Get).Methods("GET")
	r.HandleFunc("/courses/{id}", course.Update).Methods("PATCH")
	r.HandleFunc("/courses/{id}", course.Delete).Methods("DELETE")
}
