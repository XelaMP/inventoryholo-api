package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/XelaMP/inventoryholo-api/models"
	"github.com/XelaMP/inventoryholo-api/query"
	"strconv"
	"time"
)

type MovementDB struct {
	Ctx   string
	Query models.QueryDB
}

func (db MovementDB) GetAllWarehouseFilter(filter models.Filter) ([]models.Movement, error) {
	res := make([]models.Movement, 0)

	tsql := fmt.Sprintf(db.Query["listWarehouseFilter"].Q, filter.ID, filter.Type, filter.DateFrom, filter.DateTo)
	rows, err := DB.Query(tsql)
	err = db.scan(rows, err, &res, db.Ctx, "GetAllWarehouseFilter")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, err

}

func (db MovementDB) GetAllWarehouse(idWarehouse string) ([]models.Movement, error) {
	res := make([]models.Movement, 0)

	tsql := fmt.Sprintf(db.Query["listWarehouseId"].Q, idWarehouse)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAllWarehouse")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, err
}
func (db MovementDB) GetAll() ([]models.Movement, error) {
	res := make([]models.Movement, 0)

	tsql := fmt.Sprintf(db.Query["list"].Q)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return res, err
	}
	defer rows.Close()
	return res, nil

}

func (db MovementDB) Get(id string) (models.Movement, error) {
	res := make([]models.Movement, 0)

	tsql := fmt.Sprintf(db.Query["get"].Q, id)
	rows, err := DB.Query(tsql)

	err = db.scan(rows, err, &res, db.Ctx, "GetAll")
	if err != nil {
		return models.Movement{}, err
	}
	defer rows.Close()
	return res[0], nil
}

func (db MovementDB) Create(item models.Movement) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["insert"].Q)
	date, err := time.Parse(time.RFC3339, item.Date+"T05:00:00Z")
	checkError(err, "Create", db.Ctx, "Convert Date")

	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("IdProduct", item.IdProduct),
		sql.Named("IdWarehouse", item.IdWarehouse),
		sql.Named("IdUser", item.IdUser),
		sql.Named("IdClient", item.IdClient),
		sql.Named("Date", date),
		sql.Named("Quantity", item.Quantity),
		sql.Named("Type", item.Type))

	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db MovementDB) Update(id string, item models.Movement) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["update"].Q)
	date, err := time.Parse(time.RFC3339, item.Date+"T05:00:00Z")
	checkError(err, "Update", db.Ctx, "Convert Date")

	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id),
		sql.Named("IdProduct", item.IdProduct),
		sql.Named("IdWarehouse", item.IdWarehouse),
		sql.Named("IdUser", item.IdUser),
		sql.Named("IdClient", item.IdClient),
		sql.Named("Date", date),
		sql.Named("Quantity", item.Quantity),
		sql.Named("Type", item.Type))

	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func (db MovementDB) Delete(id string) (int64, error) {
	ctx := context.Background()
	tsql := fmt.Sprintf(db.Query["delete"].Q)
	result, err := DB.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id))
	if err != nil {
		return -1, err
	}
	return result.RowsAffected()
}

func GetStock(idWarehouse string, idProduct int) float64 {
	var item float64

	tsql := fmt.Sprintf(query.Movement["stock"].Q, idWarehouse, idProduct)
	rows, err := DB.Query(tsql)

	if err != nil {
		checkError(err, "GetStock", "Movement DB", "Reading rows")
		return 0
	}
	for rows.Next() {
		var stock sql.NullFloat64
		err := rows.Scan(&stock)
		item = stock.Float64
		if err != nil {
			checkError(err, "GetStock", "Movement DB", "Scan rows")
			return 0
		}
	}
	defer rows.Close()
	return item
}

func (db MovementDB) scan(rows *sql.Rows, err error, res *[]models.Movement, ctx string, situation string) error {
	var item models.Movement
	if err != nil {
		checkError(err, situation, ctx, "Reading rows")
		return err
	}

	for rows.Next() {
		err := rows.Scan(&item.ID, &item.IdProduct, &item.IdWarehouse, &item.IdUser,
			&item.IdClient, &item.Date, &item.Quantity, &item.Type)
		if err != nil {
			checkError(err, situation, ctx, "Scan rows")
			return err
		} else {
			product, _ := ProductDB{
				Ctx:   "Movement",
				Query: query.Product,
			}.Get(strconv.Itoa(item.IdProduct))
			item.Product = product.Name
			*res = append(*res, item)
		}
	}
	return nil
}
