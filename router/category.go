package router

import (
	category "github.com/XelaMP/inventoryholo-api/controllers"
	mid "github.com/XelaMP/inventoryholo-api/middleware"
	"github.com/gorilla/mux"
)

func categoryRoutes(s *mux.Router) {
	s.HandleFunc("/", category.GetCategorys).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(category.GetCategory)).Methods("GET")
	s.HandleFunc("/", category.CreateCategory).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(category.UpdateCategory)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(category.DeleteCategory)).Methods("DELETE")
}
