package Items

import (
	"encoding/json"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/jensskott/supermarket-api/Cassandra"
)

// GetOne function to list one item
func GetOne(w http.ResponseWriter, r *http.Request) {
	var item Item
	var errs []string
	var found bool = false

	vars := mux.Vars(r)
	id := vars["user_uuid"]

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

	if found {
		json.NewEncoder(w).Encode(GetItemResponse{Item: item})
	} else {
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
