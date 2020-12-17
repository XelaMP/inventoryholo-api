package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/XelaMP/inventoryholo-api/constants"
	"github.com/XelaMP/inventoryholo-api/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
)

func GetSystemUsers() []models.SystemUser {
	res := make([]models.SystemUser, 0)
	var item models.SystemUser

	tsql := fmt.Sprintf(QuerySystemUser["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next() {
		var idWarehouse sql.NullString
		err := rows.Scan(&item.ID, &item.Username, &item.Password, &item.Rol, &item.IdPerson, &idWarehouse)
		item.IdWarehouse = -1
		if idWarehouse.Valid {
			item.IdWarehouse, _ = strconv.Atoi(idWarehouse.String)
		}
		if err != nil {
			log.Println(err)
			return res
		} else {
			res = append(res, item)
		}
	}
	defer rows.Close()
	return res
}

func GetSystemUser(id string) []models.SystemUser {
	res := make([]models.SystemUser, 0)
	var item models.SystemUser

	tsql := fmt.Sprintf(QuerySystemUser["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows 1: " + err.Error())
		return res
	}
	for rows.Next() {
		var idWarehouse sql.NullString
		err := rows.Scan(&item.ID, &item.Username, &item.Password, &item.Rol, &item.IdPerson, &idWarehouse)
		item.IdWarehouse = -1
		if idWarehouse.Valid {
			item.IdWarehouse, _ = strconv.Atoi(idWarehouse.String)
		}
		if err != nil {
			log.Println(err)
			return res
		} else {
			res = append(res, item)
		}
	}
	defer rows.Close()
	return res
}

func CreateSystemUser(item models.SystemUser, noWare bool) (int64, error) {
	ctx := context.Background()

	tsql := fmt.Sprintf(QuerySystemUser["insert"].Q)
	if noWare {
		tsql = fmt.Sprintf(QuerySystemUser["insertNoWare"].Q)
	}
	item.Password = encrypt(item.Password)

	var err error
	var result sql.Result
	if !noWare {
		result, err = DB.ExecContext(
			ctx,
			tsql,
			sql.Named("UserName", item.Username),
			sql.Named("Password", item.Password),
			sql.Named("Rol", item.Rol),
			sql.Named("IdPerson", item.IdPerson),
			sql.Named("IdWarehouse", item.IdWarehouse))
	} else {
		result, err = DB.ExecContext(
			ctx,
			tsql,
			sql.Named("UserName", item.Username),
			sql.Named("Password", item.Password),
			sql.Named("Rol", item.Rol),
			sql.Named("IdPerson", item.IdPerson))
	}


	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func UpdateSystemUser(item models.SystemUser, noWare bool) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(QuerySystemUser["update"].Q)

	if noWare {
		tsql = fmt.Sprintf(QuerySystemUser["updateNoWare"].Q)
	}

	user := GetSystemUser(strconv.Itoa(item.ID))[0]
	if user.Password != item.Password {
		item.Password = encrypt(item.Password)
	}
	var err error
	var result sql.Result

	if !noWare {
		result, err = DB.ExecContext(
			ctx,
			tsql,
			sql.Named("ID", item.ID),
			sql.Named("UserName", item.Username),
			sql.Named("Password", item.Password),
			sql.Named("Rol", item.Rol),
			sql.Named("IdPerson", item.IdPerson),
			sql.Named("IdWarehouse", item.IdWarehouse))
	} else {
		result, err = DB.ExecContext(
			ctx,
			tsql,
			sql.Named("ID", item.ID),
			sql.Named("UserName", item.Username),
			sql.Named("Password", item.Password),
			sql.Named("Rol", item.Rol),
			sql.Named("IdPerson", item.IdPerson))
	}


	if err != nil {
		log.Println(err)
		return -1, err
	}
	return result.RowsAffected()
}
func DeleteSystemUser(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(QuerySystemUser["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
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
	for rows.Next() {
		var idWarehouse sql.NullString
		err := rows.Scan(&item.ID, &item.Username, &item.Password, &item.Rol, &item.IdPerson, &idWarehouse)
		item.IdWarehouse = -1
		if idWarehouse.Valid {
			item.IdWarehouse, _ = strconv.Atoi(idWarehouse.String)
		}
		if err != nil {
			log.Println(err)
			return res
		} else {
			res = append(res, item)
		}
	}
	defer rows.Close()
	return res
}

func ValidateSystemUserLogin(user string, password string) (constants.State, string) {
	items := GetSystemUserFromUserName(user)
	if len(items) > 0 {
		if comparePassword(items[0].Password, password) {
			return constants.Accept, strconv.Itoa(items[0].ID)
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
