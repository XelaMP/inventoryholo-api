package router

import (
	"github.com/XelaMP/inventoryholo-api/controllers"
	"github.com/XelaMP/inventoryholo-api/db"
	mid "github.com/XelaMP/inventoryholo-api/middleware"
	"github.com/XelaMP/inventoryholo-api/query"
	"github.com/gorilla/mux"
)

func categoryRoutes(s *mux.Router) {
	ctrl := controllers.CategoryController{
		DB: db.CategoryDB{
			Ctx:   "Category DB",
			Query: query.Category,
		},
	}
	s.HandleFunc("/", mid.CheckSecurity(ctrl.GetAll)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Get)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.Create)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Update)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Delete)).Methods("DELETE")
}
