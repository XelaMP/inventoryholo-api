package router

import (
	movement "github.com/XelaMP/inventoryholo-api/controllers"
	"github.com/gorilla/mux"
)

func movementRoutes(s *mux.Router)  {
	s.HandleFunc("/all/{idWarehouse}",movement.GetMovementsWarehouse).Methods("GET")
	s.HandleFunc("/", movement.GetMovements).Methods("GET")
	s.HandleFunc("/{id}", movement.GetMovement).Methods("GET")
	s.HandleFunc("/",movement.CreateMovement).Methods("POST")
	s.HandleFunc("/{id}", movement.UpdateMovement).Methods("PUT")
	s.HandleFunc("/{id}", movement.DeleteMovement).Methods("DELETE")
	
}