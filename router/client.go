package router

import (
	client "github.com/XelaMP/inventoryholo-api/controllers"

	"github.com/gorilla/mux"
)

func clientRoutes(s *mux.Router) {
	s.HandleFunc("/", client.GetClients).Methods("GET")
	s.HandleFunc("/{id}", client.GetClient).Methods("GET")
	s.HandleFunc("/", client.CreateClient).Methods("POST")
	s.HandleFunc("/{id}", client.UpdateClient).Methods("PUT")
	s.HandleFunc("/{id}", client.DeleteClient).Methods("DELETE")
}

