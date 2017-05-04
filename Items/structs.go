package Items

import "github.com/gocql/gocql"

// Item struct for our supermarket items
type Item struct {
	ID       gocql.UUID `json:"id"`
	Name     string     `json:"name"`
	Quantity int        `json:"quantity"`
}
