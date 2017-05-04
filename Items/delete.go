package Items

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/jensskott/supermarket-api/Cassandra"
)

// Delete one item
func Delete(c *gin.Context) {
	var errs []string

	id := c.Param("id")

	deleted := false

	uuid, err := gocql.ParseUUID(id)
	if err != nil {
		errs = append(errs, err.Error())
	} else {
		query := "DELETE FROM supermarket WHERE id=? IF EXISTS;"
		err := Cassandra.Session.Query(query, uuid).Exec()
		if err != nil {
			errs = append(errs, err.Error())
		} else {
			deleted = true
		}
	}

	if deleted {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Item succesfully deleted", "resourceId": uuid})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": errs})
	}
}
