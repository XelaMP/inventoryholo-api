package router

import (
	warehouse "github.com/XelaMP/inventoryholo-api/controllers"
	"github.com/gorilla/mux"
)

func warehouseRoutes(s *mux.Router)  {
	s.HandleFunc("/", warehouse.GetWarehouses).Methods("GET")
	s.HandleFunc("/{id}", warehouse.GetWarehouse).Methods("GET")
	s.HandleFunc("/", warehouse.CreateWarehouse).Methods("POST")
	s.HandleFunc("/{id}", warehouse.UpdateWarehouse).Methods("PUT")
	s.HandleFunc("/{id}", warehouse.DeleteWarehouse).Methods("DELETE")
	
}