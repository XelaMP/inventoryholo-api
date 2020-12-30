package controllers

import (
	"encoding/json"
	"github.com/XelaMP/inventoryholo-api/db"
	"github.com/XelaMP/inventoryholo-api/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MovementController struct{
	DB db.MovementDB
}

func (m MovementController) GetAllWarehouseFilter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Filter
	_ = json.NewDecoder(r.Body).Decode(&item)

	items, err := m.DB.GetAllWarehouseFilter(item)
	if err != nil {
		returnErr(w, err, "obtener todos warehouse filter")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (m MovementController) GetAllWarehouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["idWarehouse"]

	items, err := m.DB.GetAllWarehouse(id)
	if err != nil {
		returnErr(w, err, "obtener todos warehouse")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (m MovementController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := m.DB.GetAll()
	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (m MovementController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	items, err := m.DB.Get(id)
	if err != nil {
		returnErr(w, err, "obtener")
		return
	}

	_ = json.NewEncoder(w).Encode(items)
}

func (m MovementController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Movement

	_ = json.NewDecoder(r.Body).Decode(&item)

	result, err := m.DB.Create(item)
	if err != nil {
		returnErr(w, err, "crear")
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (m MovementController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Movement
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID, _ = strconv.Atoi(id)
	result, err := m.DB.Update(id, item)
	if err != nil {
		returnErr(w, err, "actualizar")
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (m MovementController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := m.DB.Delete(id)
	if err != nil {
		returnErr(w, err, "eliminar")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

