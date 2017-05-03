package Items

import (
	"encoding/json"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/jensskott/supermarket-api/Cassandra"
)

// Get request for all items
func Get(w http.ResponseWriter, r *http.Request) {
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

	json.NewEncoder(w).Encode(AllItemsResponse{Items: itemList})
}
