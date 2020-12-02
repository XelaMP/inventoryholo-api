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

	_ = json.NewEncoder(w).Encode(items[0])
}

func GetSystemUsers (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	items := db.GetSystemUsers()
	_ = json.NewEncoder(w).Encode(items)

}

func CreateSystemUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userPerson models.UserPerson
	_ = json.NewDecoder(r.Body).Decode(&userPerson)
	person := models.Person{
		Name: userPerson.Name,
		LastName: userPerson.LastName,
		Adress: userPerson.Adress,
		Mail: userPerson.Mail,
		Phone: userPerson.Phone,
		Dni: userPerson.Dni,
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
	var item models.SystemUser
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID, _ = strconv.Atoi(id)
	result, err := db.UpdateSystemUser(item)
	if err != nil {
		log.Println(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}

func DeleteSystemUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := db.DeleteSystemUser(id)
	if err != nil {
		log.Println(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}

