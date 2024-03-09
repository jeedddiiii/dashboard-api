package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"

	"fmt"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	TransactionID  int       `db:"transaction_id"`
	Name           string    `db:"name"`
	DateTime       time.Time `db:"date_time"`
	SourceID       int       `db:"source_id"`
	Emotion        string    `db:"emotion"`
	FaceImg        string    `db:"face_img"`
	EnvironmentImg string    `db:"environment_img"`
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	router.Use(cors.New(config))

	connectDb()
	// Define a route
	router.GET("/transaction", func(c *gin.Context) {

		page, _ := strconv.Atoi(c.Query("page"))
		limit, _ := strconv.Atoi(c.Query("limit"))

		// Set default values if not provided
		if page <= 0 {
			page = 1
		}
		if limit <= 0 {
			limit = 10 // set your default limit
		}

		// Calculate offset based on page and limit
		offset := (page - 1) * limit
		fmt.Printf("Page: %d, Limit: %d, Offset: %d\n", page, limit, offset)

		var totalCount int
		err := db.QueryRow("SELECT COUNT(*) FROM transactions").Scan(&totalCount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rows, err := db.Query("SELECT * FROM transactions LIMIT $1 OFFSET $2", limit, offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var transactions []Transaction // assuming Transaction is a struct that matches the table schema
		for rows.Next() {
			var t Transaction
			err := rows.Scan(&t.TransactionID, &t.Name, &t.DateTime, &t.Emotion, &t.SourceID, &t.FaceImg, &t.EnvironmentImg)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			transactions = append(transactions, t)
		}

		c.JSON(http.StatusOK, gin.H{
			"transactions": transactions,
			"totalCount":   totalCount,
		})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
