package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()
	connectDb()
	// Define a route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, this is your Gin server!",
		})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
