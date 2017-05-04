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
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	item := Item{Name: c.PostForm("name"), Quantity: quantity}

	gocqlUUID = gocql.TimeUUID()

	// write data to Cassandra
	err := Cassandra.Session.Query(`
		      	  INSERT INTO supermarket (id, name, quantity) VALUES (?, ?, ?)`, gocqlUUID, &item.Name, &item.Quantity).Exec()
	if err != nil {
		errs = append(errs, err.Error())
	} else {
		created = true
	}

	if created {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Item created successfully!", "resourceId": gocqlUUID})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": errs})
	}

}
