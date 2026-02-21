package handler

import (
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"nama":    "rafi",
	})
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "alolo",
		"nama":    "mamam",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "book id is " + id,
	})
}

func QueryHandler(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "query name is " + name,
	})
}

func PostBookHandler(c *gin.Context) {
	var newBooks book.BookRequest
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