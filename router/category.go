package router

import (
	category "github.com/XelaMP/inventoryholo-api/controllers"
	"github.com/gorilla/mux"
)

func categoryRoutes(s *mux.Router) {
	s.HandleFunc("/", category.GetCategorys).Methods("GET")
	s.HandleFunc("/{id}", category.GetCategory).Methods("GET")
	s.HandleFunc("/", category.CreateCategory).Methods("POST")
	s.HandleFunc("/{id}", category.UpdateCategory).Methods("PUT")
	s.HandleFunc("/{id}", category.DeleteCategory).Methods("DELETE")
}
