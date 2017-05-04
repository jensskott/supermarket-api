package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jensskott/supermarket-api/Cassandra"
	"github.com/jensskott/supermarket-api/Items"
)

func main() {

	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()

	// For release mode
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Main API router
	v1 := router.Group("/api/v1/items")
	{
		v1.POST("/", Items.Post)
		v1.GET("/", Items.Get)
		v1.GET("/:id", Items.GetOne)
		v1.PUT("/:id", Items.Update)
		v1.DELETE("/:id", Items.Delete)
	}

	// Healthcheck
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"supermarketapiv1": "running",
		})
	})
	// Run on :8080
	router.Run()
}
