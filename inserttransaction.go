package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertTransaction(c *gin.Context) {
	var t Transaction

	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO transactions (name, date_time, emotion, source_id, face_img, environment_img) VALUES ($1, $2, $3, $4, $5, $6)",
		t.Name, t.DateTime, t.Emotion, t.SourceID, t.FaceImg, t.EnvironmentImg)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": "Transaction inserted successfully"})
}
