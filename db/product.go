package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/XelaMP/inventoryholo-api/models"
	"log"
)

func GetProducts() [] models.Product {
	res := make([]models.Product, 0)
	var item models.Product

	tsql := fmt.Sprintf(QueryProduct["list"].Q)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Stock)
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

func GetProduct(id string) []models.Product {
	res := make([]models.Product, 0)
	var item models.Product

	tsql := fmt.Sprintf(QueryProduct["get"].Q, id)
	rows, err := DB.Query(tsql)

	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return res
	}
	for rows.Next(){
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Stock, &item.IdCategory)
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

func CreateProduct(item models.Product) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(QueryProduct["insert"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("Name", item.Name),
		sql.Named("description", item.Description),
		sql.Named("Price", item.Price),
		sql.Named("Stock", item.Stock),
		sql.Named("IdCategory",item.IdCategory))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
func UpdateProduct(item models.Product) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(QueryProduct["update"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", item.ID),
		sql.Named("Name", item.Name),
		sql.Named("description",item.Description),
		sql.Named("Price", item.Price),
		sql.Named("Stock",item.Stock),
		sql.Named("IdCategory",item.IdCategory))

	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
func DeleteProduct(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(QueryProduct["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}
