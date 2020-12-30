package query

import "github.com/XelaMP/inventoryholo-api/models"

var category = models.TableDB{
	Name:   "dbo.Category",
	Fields: []string{"IdCategory", "Name"},
}

var Category = models.QueryDB{
	"get":    {Q: "select " + fieldString(category.Fields) + " from " + category.Name + " where " + category.Fields[0] + " = '%s';"},
	"list":   {Q: "select " + fieldString(category.Fields) + " from " + category.Name + ";"},
	"insert": {Q: "insert into " + category.Name + " (" + fieldStringInsert(category.Fields) + ") values (" + valuesString(category.Fields) + ");"},
	"update": {Q: "update " + category.Name + " set " + updatesString(category.Fields) + " where " + category.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + category.Name + " where " + category.Fields[0] + " = @ID"},
}
