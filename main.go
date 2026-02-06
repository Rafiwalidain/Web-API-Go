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

	router.GET("/book/:id", booksHandler)

	router.GET("/query", queryHandler)

	router.POST("/book", postBookHandler)

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

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "book id is " + id,
	})
}

func queryHandler(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "query name is " + name,
	})
}

type Books struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func postBookHandler(c *gin.Context) {
	var newBooks Books
	if err := c.ShouldBindJSON(&newBooks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"title":  newBooks.Title,
			"author": newBooks.Author,
		})
	}
}
