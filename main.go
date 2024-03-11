package main

import (
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	router.Use(cors.New(config))

	connectDb()
	router.GET("/transaction", GetTransaction)
	router.GET("summary", GetSummary)

	router.Run(":8080")
}
