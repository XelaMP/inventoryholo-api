package router

import (
	user "github.com/XelaMP/inventoryholo-api/controllers"
	"github.com/gorilla/mux"
)

func userRoutes(s *mux.Router) {
	s.HandleFunc("/",user.GetSystemUsers).Methods("GET")
	s.HandleFunc("/{id}", user.GetSystemUser).Methods("GET")

	//s.HandleFunc("/{id}", mid.CheckSecurity(user.UpdatePasswordSystemUser)).Methods("PUT")
	s.HandleFunc("/", user.CreateSystemUser).Methods("POST")
	s.HandleFunc("/{id}", user.UpdateSystemUser).Methods("PUT")
	s.HandleFunc("/{id}", user.DeleteSystemUser).Methods("DELETE")
}
