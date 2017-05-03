package Items

import (
	"github.com/gocql/gocql"
)

// Item struct for our supermarket items
type Item struct {
	ID       gocql.UUID `json:"id"`
	Name     string     `json:"name"`
	Quantity int        `json:"quantity"`
}

// GetItemResponse to form payload returning a single Item struct
type GetItemResponse struct {
	Item Item `json:"item"`
}

// AllItemsResponse to form payload of an array of Item structs
type AllItemsResponse struct {
	Items []Item `json:"items"`
}

// NewItemResponse builds a payload of new user resource ID
type NewItemResponse struct {
	ID gocql.UUID `json:"id"`
}

// ErrorResponse returns an array of error strings if appropriate
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
