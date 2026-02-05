package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// This is the main function where the program execution starts.

	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/ping", pingHandler)

	router.GET("/book/:id")

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"nama":    "rafi",
	})
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "alolo",
		"nama":    "mamam",
	})
}
