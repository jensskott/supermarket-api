package Items

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/jensskott/supermarket-api/Cassandra"
)

// Update item
func Update(c *gin.Context) {
	var errs []string

	id := c.Param("id")

	updated := false

	// Convert post to int
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	item := Item{Name: c.PostForm("name"), Quantity: quantity}

	uuid, err := gocql.ParseUUID(id)
	if err != nil {
		errs = append(errs, err.Error())
	} else {
		query := "UPDATE supermarket SET name = ?, quantity = ? WHERE id = ?"
		err := Cassandra.Session.Query(query, &item.Name, &item.Quantity, uuid).Exec()
		if err != nil {
			errs = append(errs, err.Error())
		} else {
			updated = true
		}

	}

	// Return 200OK or errors
	if updated {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Item created successfully!", "resourceId": uuid})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": errs})
	}
}
