package route

import (
	"test/controller"
	"test/route/api"

	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/", controller.ServeHome).Methods("GET")

	apiRouter := r.PathPrefix("/api").Subrouter()
	api.AddApiRoutes(apiRouter)
}
