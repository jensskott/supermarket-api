package Items

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/jensskott/supermarket-api/Cassandra"
)

// Post fucntion to submit data into cassandra
func Post(c *gin.Context) {
	var errs []string
	var gocqlUUID gocql.UUID

	created := false

	// Convert post to int
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	item := Item{Name: c.PostForm("name"), Quantity: quantity}

	gocqlUUID = gocql.TimeUUID()

	// Update data in Cassandra
	err := Cassandra.Session.Query(`
		      	  INSERT INTO supermarket (id, name, quantity) VALUES (?, ?, ?)`, gocqlUUID, &item.Name, &item.Quantity).Exec()
	if err != nil {
		errs = append(errs, err.Error())
	} else {
		created = true
	}

	// Return 200OK or error
	if created {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Item created successfully!", "resourceId": gocqlUUID})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": errs})
	}

}
