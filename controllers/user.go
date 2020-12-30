package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/XelaMP/inventoryholo-api/db"
	"github.com/XelaMP/inventoryholo-api/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserController struct {
	 DB db.UserDB
	 PersonDB db.PersonDB
}

func (u UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make([]models.UserPerson, 0)
	users, err := u.DB.GetAll()
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al obtener todos, error: %v", err))
		return
	}

	for _, items := range users {
		person, _ := u.PersonDB.Get(strconv.Itoa(int(items.IdPerson)))
		userPerson := models.UserPerson{
			ID:          items.ID,
			PersonID:    int64(person.ID),
			Username:    items.Username,
			Password:    items.Password,
			Rol:         items.Rol,
			Name:        person.Name,
			LastName:    person.LastName,
			Cel:         person.Cel,
			Phone:       person.Phone,
			Address:     person.Address,
			Dni:         person.Dni,
			Mail:        person.Mail,
			IdWarehouse: items.IdWarehouse,

	}
		res = append(res, userPerson)
	}
	_ = json.NewEncoder(w).Encode(res)
}

func (u UserController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	e, err := u.DB.Get(id)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al obtener, error: %v", err))
		return
	}


		person, _ := u.PersonDB.Get(strconv.Itoa(int(e.IdPerson)))
		userPerson := models.UserPerson{
			ID:          e.ID,
			PersonID:    int64(person.ID),
			Cel:         person.Cel,
			Username:    e.Username,
			Password:    e.Password,
			Rol:         e.Rol,
			Name:        person.Name,
			LastName:    person.LastName,
			Phone:       person.Phone,
			Address:     person.Address,
			Dni:         person.Dni,
			Mail:        person.Mail,
			IdWarehouse: e.IdWarehouse,
		}

	_ = json.NewEncoder(w).Encode(userPerson)

}

func (u UserController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userPerson models.UserPerson
	_ = json.NewDecoder(r.Body).Decode(&userPerson)
	person := models.Person{
		Name:     userPerson.Name,
		LastName: userPerson.LastName,
		Cel:      userPerson.Cel,
		Address:  userPerson.Address,
		Mail:     userPerson.Mail,
		Phone:    userPerson.Phone,
		Dni:      userPerson.Dni,
	}
	idPerson, err := u.PersonDB.Create(person)

	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al crear Person, error: %v", err))
		return
	}
	user := models.SystemUser{
		Username: userPerson.Username,
		Password: userPerson.Password,
		Rol:      userPerson.Rol,
		IdPerson: idPerson,
		IdWarehouse: userPerson.IdWarehouse,
	}
	result, err := u.DB.Create(user)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al crear, error: %v", err))
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (u UserController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, _ := params["id"]

	var item models.UserPerson
	_ = json.NewDecoder(r.Body).Decode(&item)

	person := models.Person{
		ID:       int(item.PersonID),
		Name:     item.Name,
		LastName: item.LastName,
		Cel:      item.Cel,
		Phone:    item.Phone,
		Address:  item.Address,
		Dni:      item.Dni,
		Mail:     item.Mail,
	}
	result, err := u.PersonDB.Update(strconv.Itoa(int(item.PersonID)), person)

	user := models.SystemUser{
		ID:       item.ID,
		Username: item.Username,
		Password: item.Password,
		Rol:      item.Rol,
		IdPerson: item.PersonID,
		IdWarehouse: item.IdWarehouse,
	}
	result, err = u.DB.Update(id, user)
	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al actualizar, error: %v", err))
		return
	}
	_ = json.NewEncoder(w).Encode(result)

}

func (u UserController) Delete (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	user, _ := u.DB.Get(id)
	result, err := u.PersonDB.Delete(strconv.Itoa(int(user.IdPerson)))
	result, err =  u.DB.Delete(id)

	if err != nil {
		_, _ = fmt.Fprintln(w, fmt.Sprintf("Hubo un error al eliminar, error: %v", err))
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}
