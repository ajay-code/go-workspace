package router

import (
	"mongo/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/movies", controller.GetAllMyMovies).Methods("GET")
	api.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	api.HandleFunc("/movies/{id}", controller.MarkWatched).Methods("PATCH")
	api.HandleFunc("/movies/{id}", controller.DeleteMovieByID).Methods("DELETE")
	api.HandleFunc("/deleteallmovies", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
