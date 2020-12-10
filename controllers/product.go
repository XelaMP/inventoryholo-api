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

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := db.GetProducts()
	_ = json.NewEncoder(w).Encode(items)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]

		items := db.GetProduct(id)
		var product models.Product
	if len(items) > 0 {
			product = models.Product{
				ID:          items[0].ID,
				Name:        items[0].Name,
				Description: items[0].Description,
				Price:       items[0].Price,
				Stock:       items[0].Stock,
				IdCategory:  items[0].IdCategory,
			}
	}
	_ = json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item models.Product
	_ = json.NewDecoder(r.Body).Decode(&item)
	result, err := db.CreateProduct(item)
	if err != nil {
		log.Println(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	var item models.Product
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID, _ = strconv.Atoi(id)
	result, err := db.UpdateProduct(item)
	if err != nil {
		log.Println(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, _ := params["id"]
	result, err := db.DeleteProduct(id)
	if err != nil {
		log.Println(err)
	}

	_ = json.NewEncoder(w).Encode(result)
}

