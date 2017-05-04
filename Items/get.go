package Items

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/jensskott/supermarket-api/Cassandra"
)

// Get all items
func Get(c *gin.Context) {
	var itemList []Item
	m := map[string]interface{}{}

	query := "SELECT id,quantity,name FROM supermarket"
	iterable := Cassandra.Session.Query(query).Iter()
	for iterable.MapScan(m) {
		itemList = append(itemList, Item{
			ID:       m["id"].(gocql.UUID),
			Quantity: m["quantity"].(int),
			Name:     m["name"].(string),
		})
		m = map[string]interface{}{}
	}

	// Return status not found if there is no items in list
	if len(itemList) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Items found!"})
		return
	}

	// Return map of items
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": itemList})
}
