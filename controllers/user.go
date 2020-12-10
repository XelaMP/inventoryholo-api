package controllers

import (
	"encoding/json"
	"github.com/XelaMP/inventoryholo-api/db"
	"github.com/XelaMP/inventoryholo-api/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetSystemUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := db.GetSystemUser(id)
	var userPerson models.UserPerson
	if len(items) > 0 {
		person := db.GetPerson(strconv.Itoa(int(items[0].IdPerson)))
		userPerson = models.UserPerson{
			ID:       items[0].ID,
			PersonID: int64(person[0].ID),
			Username: items[0].Username,
			Password: items[0].Password,
			Rol:      items[0].Rol,
			Name:     person[0].Name,
			LastName: person[0].LastName,
			Cel:      person[0].Cel,
			Phone:    person[0].Phone,
			Address:   person[0].Address,
			Dni:      person[0].Dni,
			Mail:     person[0].Mail,
		}
	}
	_ = json.NewEncoder(w).Encode(userPerson)
}

func GetSystemUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var itemperson models.UserPerson

	//_ = json.NewDecoder(r.Body).Decode(&itemperson)
	itemperson.ID, _ = strconv.Atoi(id)

	users := db.GetSystemUsers()

	if len(users) > 0 {
		itemperson = models.UserPerson{
			ID: itemperson.ID,
			Username: itemperson.Username,
			Rol: itemperson.Rol,
			Password: itemperson.Password,
		}
	}
	_ = json.NewEncoder(w).Encode(users)
	// imita la logica de obtener un user

}

func CreateSystemUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userPerson models.UserPerson
	_ = json.NewDecoder(r.Body).Decode(&userPerson)
	person := models.Person{
		Name:     userPerson.Name,
		LastName: userPerson.LastName,
		Cel:      userPerson.Cel,
		Address:   userPerson.Address,
		Mail:     userPerson.Mail,
		Phone:    userPerson.Phone,
		Dni:      userPerson.Dni,
	}
	idPerson, err := db.CreatePerson(person)
	if err != nil {
		log.Panicln(err)
	}
	user := models.SystemUser{
		Username: userPerson.Username,
		Password: userPerson.Password,
		Rol:      userPerson.Rol,
		IdPerson: idPerson,
	}
	result, err := db.CreateSystemUser(user)
	if err != nil {
		log.Println(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}
func UpdateSystemUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, _ := params["id"]

	var item models.UserPerson

	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID, _ = strconv.Atoi(id)

	person := models.Person{
		ID:       int(item.PersonID),
		Name:     item.Name,
		LastName: item.LastName,
		Cel:      item.Cel,
		Phone:    item.Phone,
		Address:   item.Address,
		Dni:      item.Dni,
		Mail:     item.Mail,
	}

	result, err := db.UpdatePerson(person)

	user := models.SystemUser{
		ID:       item.ID,
		Username: item.Username,
		Password: item.Password,
		Rol:      item.Rol,
		IdPerson: item.PersonID,
	}

	result, err = db.UpdateSystemUser(user)


	if err != nil {
		log.Println(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}

func DeleteSystemUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	user := db.GetSystemUser(id)[0]
	result, err := db.DeletePerson(strconv.Itoa(int(user.IdPerson)))
	result, err = db.DeleteSystemUser(id)
	if err != nil {
		log.Println(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}
