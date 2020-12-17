package models

type SystemUser struct {
	ID       int    `json:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
	IdPerson int64  `json:"personId"`
	IdWarehouse int  `json:"idWarehouse"`
}

type UserResult struct {
	ID   string `json:"_id"`
	Role string `json:"role"`
}

type UserLogin struct {
	User     string `json:"username"`
	Password string `json:"password"`
}


