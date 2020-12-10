package router

import (
	product "github.com/XelaMP/inventoryholo-api/controllers"
	"github.com/gorilla/mux"
	mid "github.com/XelaMP/inventoryholo-api/middleware"
)

func productRoutes(s *mux.Router)  {

	s.HandleFunc("/", mid.CheckSecurity(product.GetProducts)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(product.GetProduct)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(product.CreateProduct)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(product.UpdateProduct)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(product.DeleteProduct)).Methods("DELETE")
}
	

