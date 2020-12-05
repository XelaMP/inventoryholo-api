package router

import (
	input "github.com/XelaMP/inventoryholo-api/controllers"
	"github.com/gorilla/mux"
)

func inputRoutes(s *mux.Router)  {

	s.HandleFunc("/", input.GetInputs).Methods("GET")
	s.HandleFunc("/{id}", input.GetInput).Methods("GET")
	s.HandleFunc("/", input.CreateInput).Methods("POST")
	s.HandleFunc("/{id}", input.UpdateInput).Methods("PUT")
	s.HandleFunc("/{id}", input.DeleteInput).Methods("DELETE")
}
