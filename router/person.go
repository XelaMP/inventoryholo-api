package router

import (
	"github.com/gorilla/mux"
	person "github.com/XelaMP/inventoryholo-api/controllers"
)

func personRoutes(s *mux.Router)  {
	s.HandleFunc("/", person.GetPersons).Methods("GET")
	s.HandleFunc("/{id}", person.GetPerson).Methods("GET")
	s.HandleFunc("/", person.CreatePerson).Methods("POST")
	s.HandleFunc("/{id}", person.UpdatePerson).Methods("PUT")
	s.HandleFunc("/{id}", person.DeletePerson).Methods("DELETE")

}
