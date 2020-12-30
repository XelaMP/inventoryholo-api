package controllers

import (
	"encoding/json"
	"github.com/XelaMP/inventoryholo-api/db"
	"github.com/XelaMP/inventoryholo-api/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ClientsController struct{
	 DB db.ClientDB
	 PersonDB db.PersonDB
}

func (c ClientsController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make([]models.ClientPerson, 0)

	clients, err := c.DB.GetAll()
	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}
	for _, cli := range clients {
		person, _ := c.PersonDB.Get(strconv.Itoa(int(cli.IdPerson)))
		item := models.ClientPerson{
			ID:       cli.ID,
			PersonID: int64(person.ID),
			Type:     cli.Type,
			Name:     person.Name,
			LastName: person.LastName,
			Cel:       person.Cel,
			Phone:    person.Phone,
			Address:  person.Address,
			Dni:      person.Dni,
			Mail:     person.Mail,
		}
		res = append(res, item)
	}
	_ = json.NewEncoder(w).Encode(res)
}

func (c ClientsController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items, err := c.DB.Get(id)
	if err != nil {
		returnErr(w, err, "obtener")
		return
	}
		person, _ := c.PersonDB.Get(strconv.Itoa(int(items.IdPerson)))
		userPerson := models.ClientPerson{
			ID:       items.ID,
			PersonID: int64(person.ID),
			Cel:      person.Cel,
			Type:     items.Type,
			Name:     person.Name,
			LastName: person.LastName,
			Phone:    person.Phone,
			Address:  person.Address,
			Dni:      person.Dni,
			Mail:     person.Mail,
		}

	_ = json.NewEncoder(w).Encode(userPerson)
}
func (c ClientsController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.ClientPerson
	_ = json.NewDecoder(r.Body).Decode(&item)
	person := models.Person{
		Name:     item.Name,
		LastName: item.LastName,
		Cel:      item.Cel,
		Phone:    item.Phone,
		Address:  item.Address,
		Dni:      item.Dni,
		Mail:     item.Mail,
	}
	idPerson, err := c.PersonDB.Create(person)
	if err != nil {
		returnErr(w, err, "crear person")
		return
	}
	client := models.Client{
		IdPerson: idPerson,
		Type:     item.Type,
	}
	result, err := c.DB.Create(client)
	if err != nil {
		returnErr(w, err, "crear")
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}


func (c ClientsController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, _ := params["id"]

	var item models.ClientPerson

	_ = json.NewDecoder(r.Body).Decode(&item)


	person := models.Person{
		ID:       int(item.PersonID),
		Name:     item.Name,
		LastName: item.LastName,
		Phone:    item.Phone,
		Cel:      item.Cel,
		Address:  item.Address,
		Dni:      item.Dni,
		Mail:     item.Mail,
	}

	result, err := c.PersonDB.Update(strconv.Itoa(int(item.PersonID)), person)

	if err != nil {
		returnErr(w, err, "actualizar person")
		return
	}

	client := models.Client{
		ID:       item.ID,
		Type:     item.Type,
		IdPerson: item.PersonID,
	}

	result, err = c.DB.Update(id, client)

	if err != nil {
		returnErr(w, err, "actualizar")
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (c ClientsController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	client, _ := c.DB.Get(id)
	result, err := c.PersonDB.Delete(strconv.Itoa(int(client.IdPerson)))
	if err != nil {
		returnErr(w, err, "eliminar person")
		return
	}
	result, err = c.DB.Delete(id)
	if err != nil {
		returnErr(w, err, "eliminar")
		return
	}


	_ = json.NewEncoder(w).Encode(result)
}
