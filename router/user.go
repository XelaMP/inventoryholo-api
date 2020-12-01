package router

import (
	user "github.com/XelaMP/inventoryholo-api/controllers"
	mid "github.com/XelaMP/inventoryholo-api/middleware"
	"github.com/gorilla/mux"
)

func userRoutes(s *mux.Router) {
	// s.HandleFunc("/", user.GetUsers).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(user.GetSystemUser)).Methods("GET")
	// s.HandleFunc("/{id}", mid.CheckSecurity(user.UpdatePasswordSystemUser)).Methods("PUT")
	// s.HandleFunc("/", user.CreateUser).Methods("POST")
	// s.HandleFunc("/{id}", mid.CheckSecurity(user.UpdateUser)).Methods("PUT")
	// s.HandleFunc("/{id}", mid.CheckSecurity(user.DeleteUser)).Methods("DELETE")
}

