package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	var exists bool
	err = db.QueryRow("SELECT exists (SELECT 1 FROM users WHERE name=$1)", name).Scan(&exists)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("query err: %s", err.Error()))
		return
	}

	if exists {
		c.String(http.StatusBadRequest, "User already exists")
		return
	}
	dir := filepath.Join("users", name)
	os.MkdirAll(dir, os.ModePerm)

	// Save file to directory
	dst := filepath.Join(dir, file.Filename)
	c.SaveUploadedFile(file, dst)

	// Save user to database
	_, err = db.Exec("INSERT INTO users (name, img_path) VALUES ($1, $2)", name, dst)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("insert err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, "User created successfully")

}
