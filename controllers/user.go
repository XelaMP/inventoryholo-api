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

	var item models.SystemUser
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := db.CreateSystemUser(item)
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

