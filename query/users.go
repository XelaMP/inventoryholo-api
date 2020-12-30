package query

import "github.com/XelaMP/inventoryholo-api/models"

var user = models.TableDB{
	Name:   "dbo.Users",
	Fields: []string{"IdUser", "Username", "Password", "Rol", "IdPerson", "IdWarehouse"},
}

var SystemUser = models.QueryDB{
	"getUserName": {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[1] + " = '%s';"},
	"get":         {Q: "select " + fieldString(user.Fields) + " from " + user.Name + " where " + user.Fields[0] + " = '%s';"},
	"list":        {Q: "select " + fieldString(user.Fields) + " from " + user.Name + ";"},
	//"getidPerson" : {Q: "select" + fieldString(user.Fields)+ "from" + person.Name + "where" + person.Fields[0] + "= %s;"},
	"insert":       {Q: "insert into " + user.Name + "(" + fieldStringInsert(user.Fields) + ") values (" + valuesString(user.Fields) + ");"},
	//"insertNoWare": {Q: "insert into " + user.Name + "(" + fieldStringInsert(user.Fields[:5]) + ") values (" + valuesString(user.Fields[:5]) + ");"},
	"update":       {Q: "update " + user.Name + " set " + updatesString(user.Fields) + " where " + user.Fields[0] + " = @ID;"},
	//"updateNoWare": {Q: "update " + user.Name + " set " + updatesString(user.Fields[:5]) + " where " + user.Fields[0] + " = @ID;"},
	"delete":       {Q: "delete from " + user.Name + " where " + user.Fields[0] + " = @ID"},
}

