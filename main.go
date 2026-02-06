package main

import (
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// This is the main function where the program execution starts.

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)

	v1.GET("/ping", handler.PingHandler)

	v1.GET("/book/:id", handler.BooksHandler)

	v1.GET("/query", handler.QueryHandler)

	v1.POST("/book", handler.PostBookHandler)

	router.Run()
}
