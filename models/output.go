package models

type Output struct {
	IdOutput int `json:"id_output"`
	Fecha string `json:"fecha"`
	Quantity int `json:"quantity"`
	IdProduct int `json:"id_product"`
}
