package book

// import "gorm.io/gorm"


type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	book := Book{
		Title:  bookRequest.Title,
		Author: bookRequest.Author,
		Price:  bookRequest.Price,
		Rating: bookRequest.Rating,
	}
	return s.repository.Create(book)
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	book := Book{
		Title:  bookRequest.Title,
		Author: bookRequest.Author,
		Price:  bookRequest.Price,
		Rating: bookRequest.Rating,
	}
	return s.repository.Update(book)
}

func (s *service) Delete(ID int) error {
	return s.repository.Delete(ID)
}