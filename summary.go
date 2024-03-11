package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetSummary(c *gin.Context) {

	rows, err := db.Query(`
        SELECT name, emotion, COUNT(*) as count
        FROM transactions
        WHERE name != 'Unknown'
        GROUP BY name, emotion
    `)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	results := make(map[string]map[string]int)
	for rows.Next() {
		var name, emotion string
		var count int
		err := rows.Scan(&name, &emotion, &count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if _, ok := results[name]; !ok {
			results[name] = make(map[string]int)
		}
		results[name][emotion] = count
	}

	c.JSON(http.StatusOK, results)
}
