package models

type UserPerson struct {
	ID          int    `json:"_id"`
	PersonID    int64  `json:"idPerson"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Rol         string `json:"rol"`
	Name        string `json:"name"`
	LastName    string `json:"lastName"`
	Cel         string `json:"cel"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Dni         string `json:"dni"`
	Mail        string `json:"mail"`
	IdWarehouse int    `json:"idWarehouse"`
}

type ClientPerson struct {
	ID       int    `json:"_id"`
	Type     string `json:"type"`
	PersonID int64  `json:"personId"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Cel      string `json:"cel"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Dni      string `json:"dni"`
	Mail     string `json:"mail"`
}
