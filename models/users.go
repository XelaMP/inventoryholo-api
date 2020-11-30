package models

type SystemUser struct {
	IdUser       int `json:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	rol string    `json:"typeUser"`
	IdPerson int `json:"personId"`
}

type UserLogin struct {
	User string
	Password string


}


