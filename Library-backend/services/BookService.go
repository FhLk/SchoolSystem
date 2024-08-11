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

type BookServiceImpl struct {
	bookRepository repositories.BookRepository
}

func NewBookService(bookRepository repositories.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{bookRepository: bookRepository}
}

func (s *BookServiceImpl) CreateBook(ctx context.Context, book *models.Book) error {
	err := s.bookRepository.CreateBook(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookServiceImpl) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	book, err := s.bookRepository.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookServiceImpl) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	books, err := s.bookRepository.GetAllBooks(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookServiceImpl) DeleteBook(ctx context.Context, id string) error {
	err := s.bookRepository.DeleteBook(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookServiceImpl) UpdateBook(ctx context.Context, book *models.Book) error {
	err := s.bookRepository.UpdateBook(ctx, book)
	if err != nil {
		return err
	}
	return nil
}
