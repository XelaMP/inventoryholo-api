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

var category = TableDB{
	Name:   "dbo.Category",
	Fields: []string{"idCategory", "Name"},
}

var queryCategory = map[string]*queryConfig{
	"get":    {Q: "select " + fieldString(category.Fields) + " from " + category.Name + " where " + category.Fields[0] + " = %s;"},
	"list":   {Q: "select " + fieldString(category.Fields) + " from " + category.Name + ";"},
	"insert": {Q: "insert into " + category.Name + " (" + fieldStringInsert(category.Fields) + ") values (" + valuesString(category.Fields) + ");"},
	"update": {Q: "update " + category.Name + " set " + updatesString(category.Fields) + " where " + category.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + category.Name + " where " + category.Fields[0] + " = @ID"},
}