package controllers

import (
	"database/sql"
	"github.com/XelaMP/inventoryholo-api/models"
)

type PersonController struct {
	Db *sql.DB
}

func (personc PersonController) FindAll() ( [] models.Person, error )  {
		rows, err := personc.Db.Query("select from Person")
	if err != nil {
		return nil, err

		
	} else {
		var persons [] models.Person
		for rows.Next() {
			var id int
			var name string
			var lastname string
			var phone int
			var adress string
			var dni int
			var mail string
			err2 := rows.Scan(&id, &name, &lastname, & phone, &adress, &dni, &mail)

			if err2 != nil {
				return nil, err2
			}

				else {
				person := models.Person{
				IdPerson: id,
				Name:     name,
				LastName: lastname,
				Phone:    phone,
				Adress:   adress,
				Dni:      dni,
				mail:     mail,
			}
			persons = append(persons, person)

		}


		}
			return persons, nil
	}
	
}

