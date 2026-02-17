package main

import (
	// "fmt"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// This is the main function where the program execution starts.

	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka-api-go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// fmt.Println("Database connected successfully")
	db.AutoMigrate(&book.Book{})

	// create book
	// book := book.Book{}
	// book.Title = "Belajar Golang ke 3"
	// book.Author = "Rafi walidan"
	// book.Price = 40000
	// book.Rating = 4

	// err = db.Create(&book).Error
	// if err != nil {
	// 	panic("failed to create book")
	// }

	var book book.Book

	// read (ada di documentasi gorm)
	err = db.Where("id = ?", 1).Find(&book).Error
	if err != nil {
		panic("failed to get book")
	}

	// update
	// book.Title = "Belajar Golang revised"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	panic("failed to update book")
	// }
	// fmt.Println("Book title:", book.Title)
	// fmt.Printf("Book objek: %+v\n", book)

	// for _, book := range books {
	// 	fmt.Println("Book title:", book.Title)
	// 	fmt.Printf("Book objek: %+v\n", book)
	// }

	// delete
	err = db.Delete(&book).Error
	if err != nil {
		panic("failed to delete book")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)

	v1.GET("/ping", handler.PingHandler)

	v1.GET("/books/:id", handler.BooksHandler)

	v1.GET("/query", handler.QueryHandler)

	v1.POST("/books", handler.PostBookHandler)

	router.Run()
}
