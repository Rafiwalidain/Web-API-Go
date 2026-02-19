package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
	books []Book
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) FindByID(id int) (Book, error) {
	var book Book
	err := r.db.Where("id = ?", id).Find(&book).Error
	return book, err
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) Delete(id int) error {
	err := r.db.Delete(&Book{}, id).Error
	return err
}