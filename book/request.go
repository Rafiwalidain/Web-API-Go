package book

type BookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Price  int    `json:"price" binding:"required"`
	Rating int    `json:"rating" binding:"required"`
}