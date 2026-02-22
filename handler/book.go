package handler

import (
	"net/http"
	"pustaka-api/book"
	// "pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService book.Service
}

func NewbookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{
		bookService: bookService,
	 }
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"nama":    "rafi",
	})
}

func (h *bookHandler) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "alolo",
		"nama":    "mamam",
	})
}

func (h *bookHandler) BooksHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "book id is " + id,
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "query name is " + name,
	})
}

func (h *bookHandler) PostBookHandler(c *gin.Context) {
	var newBooks book.BookRequest
	if err := c.ShouldBindJSON(&newBooks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 
	createBook, err := h.bookService.Create(newBooks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": createBook,
	})
}