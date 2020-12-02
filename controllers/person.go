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

func GetPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := db.GetPersons()
	_ = json.NewEncoder(w).Encode(items)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := db.GetPerson(id)

	_ = json.NewEncoder(w).Encode(items[0])
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Person
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := db.CreatePerson(item)
	if err != nil {
		log.Println(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Person
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID, _ = strconv.Atoi(id)
	result, err := db.UpdatePerson(item)
	if err != nil {
		log.Println(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := db.DeletePerson(id)
	if err != nil {
		log.Println(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}
