package services

import (
	"Library-backend/models"
	"Library-backend/repositories"
	"context"
)

type BookService interface {
	CreateBook(ctx context.Context, book *models.Book) error
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	GetAllBooks(ctx context.Context) ([]*models.Book, error)
	DeleteBook(ctx context.Context, id string) error
	UpdateBook(ctx context.Context, book *models.Book) error
}

type bookService struct {
	bookRepository repositories.BookRepository
}

func NewBookService(bookRepository repositories.BookRepository) *bookService {
	return &bookService{bookRepository: bookRepository}
}

func (s *bookService) CreateBook(ctx context.Context, book *models.Book) error {
	err := s.bookRepository.CreateBook(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (s *bookService) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	book, err := s.bookRepository.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *bookService) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	books, err := s.bookRepository.GetAllBook(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *bookService) DeleteBook(ctx context.Context, id string) error {
	err := s.bookRepository.DeleteBook(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *bookService) UpdateBook(ctx context.Context, book *models.Book) error {
	err := s.bookRepository.UpdateBook(ctx, book)
	if err != nil {
		return err
	}
	return nil
}
