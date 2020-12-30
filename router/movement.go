package router

import (
	movement "github.com/XelaMP/inventoryholo-api/controllers"
	"github.com/XelaMP/inventoryholo-api/db"
	mid "github.com/XelaMP/inventoryholo-api/middleware"
	"github.com/XelaMP/inventoryholo-api/query"
	"github.com/gorilla/mux"
)

func movementRoutes(s *mux.Router) {
	ctrl := movement.MovementController{
		DB: db.MovementDB{
			Ctx:   "Movement DB",
			Query: query.Movement,
		},
	}
	s.HandleFunc("/all/{idWarehouse}", mid.CheckSecurity(ctrl.GetAllWarehouse)).Methods("GET")
	s.HandleFunc("/filter/", mid.CheckSecurity(ctrl.GetAllWarehouseFilter)).Methods("POST")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.GetAll)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Get)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.Create)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Update)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Delete)).Methods("DELETE")

}
