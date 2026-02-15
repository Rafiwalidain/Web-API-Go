package book

import "time"

type Book struct {
	Id        int
	Title     string
	Author    string
	Price     int
	Rating    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
