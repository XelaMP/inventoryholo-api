package models

type Input struct {
	IdInput int `json:"id_input"`
	Fecha string `json:"fecha"`
	Quantity int `json:"quantity"`
	IdProduct int `json:"id_product"`
}
