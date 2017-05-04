package Cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

// Session data for cassandra
var Session *gocql.Session

// Connect to cassandra
func Connect(clusterIP string) *gocql.Session {
	var err error

	cluster := gocql.NewCluster(clusterIP)
	cluster.Keyspace = "api"
	Session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("Unable to connect to cassandra")
		panic(err)
	}
	fmt.Println("Connected to cassandra")
	return Session
}
