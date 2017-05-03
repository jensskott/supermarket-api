package Items

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/jensskott/supermarket-api/Cassandra"
)

// Post function to create new item
func Post(w http.ResponseWriter, r *http.Request) {
	var errs []string
	var gocqlUUID gocql.UUID

	item, errs := FormToItem(r)

	created := false

	if len(errs) == 0 {
		fmt.Println("creating a new item")

		gocqlUUID = gocql.TimeUUID()

		// write data to Cassandra
		err := Cassandra.Session.Query(`
      	  INSERT INTO supermarket (id, name, quantity) VALUES (?, ?, ?)`, gocqlUUID, item.Name, item.Quantity).Exec()
		if err != nil {
			errs = append(errs, err.Error())
		} else {
			created = true
		}
	}

	if created {
		fmt.Println("item_id", gocqlUUID)
		json.NewEncoder(w).Encode(NewItemResponse{ID: gocqlUUID})
	} else {
		fmt.Println("errors", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
