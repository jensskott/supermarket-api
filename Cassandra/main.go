package Cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

// Session data for cassandra
var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("34.250.35.246")
	cluster.Keyspace = "api"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to cassandra")
}
