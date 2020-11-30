package controllers

import (
	"database/sql"
	"github.com/XelaMP/inventoryholo-api"
	)

type PatientController struct {
	Db *sql.DB
}

func (pc PatientController) FindAll () (patients [] models.Person, error)  {
	rows, err := pc.Db.Query("select * from Person")
}