package routes

import (
	"github.com/ajay-code/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(r *mux.Router){

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	api.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	api.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("GET")
	api.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	api.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("Delete")

}