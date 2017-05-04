package Items

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/jensskott/supermarket-api/Cassandra"
)

// GetOne item
func GetOne(c *gin.Context) {
	var item Item
	var errs []string

	found := false

	id := c.Param("id")

	// Itterate trough cassandra requests and add to struct
	uuid, err := gocql.ParseUUID(id)
	if err != nil {
		errs = append(errs, err.Error())
	} else {
		m := map[string]interface{}{}
		query := "SELECT id,quantity,name FROM supermarket WHERE id=? LIMIT 1"
		iterable := Cassandra.Session.Query(query, uuid).Consistency(gocql.One).Iter()
		for iterable.MapScan(m) {
			found = true
			item = Item{
				ID:       m["id"].(gocql.UUID),
				Quantity: m["quantity"].(int),
				Name:     m["name"].(string),
			}
		}
		if !found {
			errs = append(errs, "Item not found")
		}
	}

	// Return item data or errors
	if found {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &item})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": errs})
		return
	}
}
