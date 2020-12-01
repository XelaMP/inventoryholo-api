package db

import "database/sql"

var DB *sql.DB

type queryConfig struct {
	Name string
	Q    string
}

type TableDB struct {
	Name   string
	Fields []string
}

var user = TableDB{
	Name:   "dbo.Users",
	Fields: []string{"idUser", "Username", "password", "rol", "idPerson"},
}

var QuerySystemUser = map[string]*queryConfig{
	"getUserName": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[2] + " = '%s';"},
	"get": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[0] + " = '%s';"},
	"updatePassword": {Q: "update " + user.Name + " set v_Password = @Password where " + user.Fields[0] + " = '%s';"},
}