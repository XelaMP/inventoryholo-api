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

type ProductController struct{
	DB db.ProductDB
}


func (p ProductController) GetAllStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	items, err := p.DB.GetAllStock(id)
	if err != nil {
		returnErr(w, err, "obtener todos")
		return
	}

	_ = json.NewEncoder(w).Encode(items)
}

func (p ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items, err := p.DB.GetAll()
	if err != nil {
		returnErr(w, err, "obtener")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (p ProductController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]


	items, err := p.DB.Get(id)
	if err != nil {
		returnErr(w, err, "obtener")
		return
	}
	_ = json.NewEncoder(w).Encode(items)
}

func (p ProductController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Product
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := p.DB.Create(item)
	if err != nil {
		returnErr(w, err, "crear")
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (p ProductController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

	var item models.Product
	_ = json.NewDecoder(r.Body).Decode(&item)

	item.ID, _ = strconv.Atoi(id)
	result, err := p.DB.Update(item)
	if err != nil {
		returnErr(w, err, "actualizar")
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func (p ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)

	id, _ := params["id"]
	result, err := p.DB.Delete(id)
	fmt.Printf("aqui el ",err)

	if err != nil {
		returnErr(w, err, "eliminar")
		return
	}

	_ = json.NewEncoder(w).Encode(result)
}
