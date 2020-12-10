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
	Fields: []string{"IdUser", "Username", "Password", "Rol", "IdPerson"},
}

var QuerySystemUser = map[string]*queryConfig{
	"getUserName": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[2] + " = %s;"},
	"get": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[0] + " = %s;"},
	"updatePassword": {Q: "update " + user.Name + " set v_Password = @Password where " + user.Fields[0] + " = %s;"},
	"list": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + ";"},
	//"getidPerson" : {Q: "select" + fieldString(user.Fields)+ "from" + person.Name + "where" + person.Fields[0] + "= %s;"},
	"insert": {Q: "insert into "+ user.Name + "("+ fieldStringInsert(user.Fields)+  ") values (" +valuesString(user.Fields) + ");"},
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

var client = TableDB{
	Name:   "dbo.Client",
	Fields: []string{"IdClient", "IdPerson", "Type"},
}

var queryClient = map[string]*queryConfig{
	"get":    {Q: "select " + fieldString(client.Fields) + " from " + client.Name + " where " + client.Fields[0] + " = %s;"},
	"list":   {Q: "select " + fieldString(client.Fields) + " from " + client.Name + ";"},
	"insert": {Q: "insert into " + client.Name + " (" + fieldStringInsert(client.Fields) + ") values (" + valuesString(client.Fields) + ");"},
	"update": {Q: "update " + client.Name + " set " + updatesString(client.Fields) + " where " + client.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + client.Name + " where " + client.Fields[0] + " = @ID"},
}



var person = TableDB{
	Name: "dbo.Person",
	Fields: []string{"IdPerson", "Name", "LastName", "Cel" ,"Phone", "Dni", "Address", "Mail"},
}

var queryPerson = map[string]*queryConfig{
	"get": {Q: "select " + fieldString(person.Fields) + " from " + person.Name + " where " + person.Fields[0] + " =%s;"},
	"list": {Q: "select " + fieldString(person.Fields) + " from " +person.Name + ";"},
	"insert": {Q: "insert into " + person.Name + "("+ fieldStringInsert(person.Fields) + ") values (" + valuesString(person.Fields) + ");"},
	"update": {Q: "update " + person.Name + " set " + updatesString(person.Fields) + " where " + person.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + person.Name + " where " + person.Fields[0] + " = @ID"},

}

var product = TableDB{
	Name: "dbo.Product",
	Fields: []string{"IdProduct", "Name" ,"description","Price","Stock","IdCategory"},
}

var QueryProduct = map[string]*queryConfig{
	"get": {Q: "select " + fieldString(product.Fields) + " from " + product.Name + " where " + product.Fields[0] + " =%s;"},
	"list": {Q: "select " + fieldString(product.Fields) + " from " +product.Name + ";"},
	"insert": {Q: "insert into " + product.Name + "("+ fieldStringInsert(product.Fields) + ") values (" + valuesString(product.Fields) + ");"},
	"update": {Q: "update " + product.Name + " set " + updatesString(product.Fields) + " where " + product.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + product.Name + " where " + product.Fields[0] + " = @ID"},

}

var movement = TableDB{
	Name:   "dbo.Movement",
	Fields: []string{"IdMovement", "IdProduct", "IdWareHouse", "IdUser", "IdClient", "Date", "Quantity","Type"},
}

var queryMovement = map[string]*queryConfig{
	"get":    {Q: "select " + fieldString(movement.Fields) + " from " + movement.Name + " where " + movement.Fields[0] + " =%s;"},
	"list":   {Q: "select " + fieldString(movement.Fields) + " from " + movement.Name + ";"},
	"listWarehouseId":   {Q: "select " + fieldString(movement.Fields) + " from " + movement.Name + " where " +
		movement.Fields[2] + " = %s;"},
	"insert": {Q: "insert into " + movement.Name + "(" + fieldStringInsert(movement.Fields) + ") values (" + valuesString(movement.Fields) + ");"},
	"update": {Q: "update " + movement.Name + " set " + updatesString(movement.Fields) + " where " + movement.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + movement.Name + " where " + movement.Fields[0] + " = @ID"},
}

var warehouse = TableDB{
	Name:   "dbo.Warehouse",
	Fields: []string{"IdWarehouse", "Name", "Address", "State"},
}

var queryWarehouse = map[string]*queryConfig{
	"get":    {Q: "select " + fieldString(warehouse.Fields) + " from " + warehouse.Name + " where " + warehouse.Fields[0] + " =%s;"},
	"list":   {Q: "select " + fieldString(warehouse.Fields) + " from " + warehouse.Name + ";"},
	"insert": {Q: "insert into " + warehouse.Name + "(" + fieldStringInsert(warehouse.Fields) + ") values (" + valuesString(warehouse.Fields) + ");"},
	"update": {Q: "update " + warehouse.Name + " set " + updatesString(warehouse.Fields) + " where " + warehouse.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + warehouse.Name + " where " + warehouse.Fields[0] + " = @ID"},
}

