package query

import "github.com/XelaMP/inventoryholo-api/models"

var movement = models.TableDB{
	Name:   "dbo.Movement",
	Fields: []string{"IdMovement", "IdProduct", "IdWarehouse", "IdUser", "IdClient", "Date", "Quantity", "Type"},
}

var Movement = models.QueryDB{
	"get":   {Q: "select " + fieldString(movement.Fields) + " from " + movement.Name + " where " + movement.Fields[0] + " = '%s';"},
	"list":  {Q: "select " + fieldString(movement.Fields) + " from " + movement.Name + ";"},
	"stock": {Q: "select sum(Quantity) as stock from Movement where IdWarehouse = %s and IdProduct = %d;"},
	"listWarehouseId": {Q: "select " + fieldString(movement.Fields) + " from " + movement.Name + " where " +
		movement.Fields[2] + " = %s;"},
	"listWarehouseFilter": {Q: "select " + fieldString(movement.Fields) + " from " + movement.Name + " where " +
		movement.Fields[2] + " = %s and " + movement.Fields[7] + " = '%s' " +
		"and CAST(Date as date) >= CAST('%s' as date) and CAST(Date as date) <= CAST('%s' as date) " +
		"order by IdMovement desc;"},
	"insert": {Q: "insert into " + movement.Name + "(" + fieldStringInsert(movement.Fields) + ") values (" + valuesString(movement.Fields) + ");"},
	"update": {Q: "update " + movement.Name + " set " + updatesString(movement.Fields) + " where " + movement.Fields[0] + " = @ID;"},
	"delete": {Q: "delete from " + movement.Name + " where " + movement.Fields[0] + " = @ID"},
}
