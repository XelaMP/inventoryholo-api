package helper

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() (db *sql.DB, err error) {
	db, err = sql.Open("mssql","server=localhost;user id=sa; password=123456")
	return

}