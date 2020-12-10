package controllers

import (
	"encoding/json"
	"github.com/XelaMP/inventoryholo-api/db"
	"github.com/XelaMP/inventoryholo-api/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetMovementsWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["idWarehouse"]

	items := db.GetMovementsWarehouse(id)
	_ = json.NewEncoder(w).Encode(items)
}

func GetMovements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := db.GetMovements()
	_ = json.NewEncoder(w).Encode(items)
}

func GetMovement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items := db.GetMovement(id)

	_ = json.NewEncoder(w).Encode(items[0])
}

func CreateMovement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Movement
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := db.CreateMovement(item)
	checkError(err, "Created", "Movement")

	_ = json.NewEncoder(w).Encode(result)
}

func UpdateMovement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Movement
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID, _ = strconv.Atoi(id)
	result, err := db.UpdateMovement(item)
	checkError(err, "Updated", "Movement")

	_ = json.NewEncoder(w).Encode(result)
}

func DeleteMovement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := db.DeleteMovement(id)
	checkError(err, "Deleted", "Movement")

	_ = json.NewEncoder(w).Encode(result)
}

