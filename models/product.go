package models

type Product struct {
	ID			int 	`json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	IdCategory  int     `json:"id_category"`
}

type ProductCategory struct {
	ProductId   int 	`json:"product_id"`
	NameProduct string 	`json:"name_product"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryId  int 	`json:"category_id"`
	CategoryName string `json:"category_name"`
}