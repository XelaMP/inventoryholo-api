package router

import "github.com/gorilla/mux"

func Routes(r *mux.Router) {
	u := r.PathPrefix("/user").Subrouter()
	userRoutes(u)
}
