package models

type Product struct {
	IdProduct   int     `json:"id_product"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	IdCategory  int     `json:"id_category"`
}
