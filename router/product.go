package router

import (
	product "github.com/XelaMP/inventoryholo-api/controllers"
	"github.com/gorilla/mux"
)

func productRoutes(s *mux.Router)  {

	s.HandleFunc("/", product.GetProducts).Methods("GET")
	s.HandleFunc("/{id}", product.GetProduct).Methods("GET")
	s.HandleFunc("/", product.CreateProduct).Methods("POST")
	s.HandleFunc("/{id}", product.UpdateProduct).Methods("PUT")
	s.HandleFunc("/{id}", product.DeleteProduct).Methods("DELETE")
}
	

