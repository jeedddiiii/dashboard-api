package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetUsers(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, img_path FROM users")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error querying database: %v", err)
		return
	}
	defer rows.Close()

	users := make([]map[string]interface{}, 0)
	for rows.Next() {
		var id int
		var name, imgPath string
		if err := rows.Scan(&id, &name, &imgPath); err != nil {
			c.String(http.StatusInternalServerError, "Error scanning row: %v", err)
			return
		}
		user := map[string]interface{}{
			"id":      id,
			"name":    name,
			"imgPath": imgPath,
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
