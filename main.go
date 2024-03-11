package main

import (
	"github.com/gin-contrib/cors"

	"log"

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

	BotInit()
	SentMessage()
	log.Println("Message sent successfully!")

	router.Run(":8080")
}
