package services

import (
	"Library-backend/models"
	"Library-backend/repositories"
)

type BookService interface {
	CreateUser(user *models.Book) error
	GetUserByID(id int) (*models.Book, error)
}

type bookService struct {
	bookRepository repositories.BookRepository
}

func NewBookService(bookRepository repositories.BookRepository) *bookService {
	return &bookService{bookRepository: bookRepository}

}
