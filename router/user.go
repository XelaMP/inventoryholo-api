package router

import (
	"github.com/XelaMP/inventoryholo-api/controllers"
	"github.com/XelaMP/inventoryholo-api/db"
	mid "github.com/XelaMP/inventoryholo-api/middleware"
	"github.com/XelaMP/inventoryholo-api/query"
	"github.com/gorilla/mux"
)

func userRoutes(s *mux.Router) {
	ctrl := controllers.UserController{
		DB: db.UserDB{
			Ctx:   "User DB",
			Query: query.SystemUser,
		},
		PersonDB: db.PersonDB{Ctx: "PersonDB", Query: query.Person},
	}
	s.HandleFunc("/", mid.CheckSecurity(ctrl.GetAll)).Methods("GET")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Get)).Methods("GET")
	s.HandleFunc("/", mid.CheckSecurity(ctrl.Create)).Methods("POST")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Update)).Methods("PUT")
	s.HandleFunc("/{id}", mid.CheckSecurity(ctrl.Delete)).Methods("DELETE")
}
