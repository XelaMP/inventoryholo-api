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
	"getUserName": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[2] + " = %s;"},
	"get": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[0] + " = %s;"},
	"updatePassword": {Q: "update " + user.Name + " set v_Password = @Password where " + user.Fields[0] + " = %s;"},
	"list": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + ";"},
	//"getidPerson" : {Q: "select" + fieldString(user.Fields)+ "from" + person.Name + "where" + person.Fields[0] + "= %s;"},
	"insert": {Q: "insert into "+ user.Name + "("+ fieldString(user.Fields)+  ") values (" +valuesString(user.Fields) + ");"},
	"update": {Q: "update " + user.Name + " set " + updatesString(user.Fields) + " where " + user.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + user.Name + " where " + user.Fields[0] + " = @ID"},


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


var person = TableDB{
	Name: "dbo.Person",
	Fields: []string{"IdPerson", "Name", "LastName", "Phone", "Dni", "Adress", "mail"},
}

var queryPerson = map[string]*queryConfig{
	"get": {Q: "select " + fieldString(person.Fields) + " from " + person.Name + " where " + person.Fields[0] + " =%s;"},
	"list": {Q: "select " + fieldString(person.Fields) + " from " +person.Name + ";"},
	"insert": {Q: "insert into " + person.Name + "("+ fieldString(person.Fields) + ") values (" + valuesString(person.Fields) + ");"},
	"update": {Q: "update " + person.Name + " set " + updatesString(person.Fields) + " where " + person.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + person.Name + " where " + person.Fields[0] + " = @ID"},

}