package db

import (
	"fmt"
	"github.com/XelaMP/inventoryholo-api/constants"
	"github.com/XelaMP/inventoryholo-api/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GetSystemUser(id string) []models.SystemUser {
	res := make([]models.SystemUser, 0)
	var item models.SystemUser

	tsql := fmt.Sprintf(QuerySystemUser["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Username, &item.Password, &item.Rol, &item.IdPerson)
		if err != nil {
			log.Println(err)
			return res
		} else{
			res = append(res, item)
		}
	}
	defer rows.Close()
	return res
}

func GetSystemUserFromUserName(userName string) []models.SystemUser {
	res := make([]models.SystemUser, 0)
	var item models.SystemUser

	tsql := fmt.Sprintf(QuerySystemUser["getUserName"].Q, userName)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Username, &item.Password, &item.Rol, &item.IdPerson)
		if err != nil {
			log.Println(err)
			return res
		} else{
			res = append(res, item)
			log.Println(item.Password)
		}
	}
	defer rows.Close()
	return res
}

func ValidateSystemUserLogin(user string, password string) (constants.State, string){
	items := GetSystemUserFromUserName(user)
	if len(items) > 0 {
		if comparePassword(items[0].Password, password) {
			return constants.Accept, string(rune(items[0].ID))
		}
		return constants.InvalidCredentials, ""
	}
	return constants.NotFound, ""
}

func comparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func encrypt(password string) string {
	passwordByte := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
		return ""
	}
	return string(hashedPassword)
}
