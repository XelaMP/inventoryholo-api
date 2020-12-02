package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/XelaMP/inventoryholo-api/models"
	"log"
)

func GetPersons() []models.Person {
	res := make([]models.Person, 0)
	var item models.Person

	tsql := fmt.Sprintf(queryPerson["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Name,&item.LastName,&item.Dni,&item.Phone,&item.Address,&item.Mail)
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
func GetPerson(id string) []models.Person {
	res := make([]models.Person, 0)
	var item models.Person

	tsql := fmt.Sprintf(queryPerson["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Name,&item.LastName,&item.Dni,&item.Phone,&item.Address,&item.Mail)
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

func CreatePerson(item models.Person) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryPerson["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", item.Name),
		sql.Named("LastName", item.LastName),
		sql.Named("Dni", item.Dni),
		sql.Named("Phone", item.Phone),
		sql.Named("Address", item.Address),
		sql.Named("Mail", item.Mail))

	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func UpdatePerson(item models.Person) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryPerson["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", item.ID),
		sql.Named("Name", item.Name),
		sql.Named("LastName", item.LastName),
		sql.Named("Dni", item.Dni),
		sql.Named("Phone", item.Phone),
		sql.Named("Address", item.Address),
		sql.Named("Mail", item.Mail))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func DeletePerson (id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(queryPerson["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}


