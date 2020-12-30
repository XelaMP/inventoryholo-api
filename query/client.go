package query

import "github.com/XelaMP/inventoryholo-api/models"

var client = models.TableDB{
	Name:   "dbo.Client",
	Fields: []string{"IdClient", "IdPerson", "Type"},
}

var Client = models.QueryDB{
	"get":    {Q: "select " + fieldString(client.Fields) + " from " + client.Name + " where " + client.Fields[0] + " = '%s';"},
	"list":   {Q: "select " + fieldString(client.Fields) + " from " + client.Name + ";"},
	"insert": {Q: "insert into " + client.Name + " (" + fieldStringInsert(client.Fields) + ") values (" + valuesString(client.Fields) + ");"},
	"update": {Q: "update " + client.Name + " set " + updatesString(client.Fields) + " where " + client.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + client.Name + " where " + client.Fields[0] + " = @ID"},
}
