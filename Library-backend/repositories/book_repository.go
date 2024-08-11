package repositories

import (
	"Library-backend/models"
	"context"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book *models.Book) error
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	GetAllBooks(ctx context.Context) ([]*models.Book, error)
	DeleteBook(ctx context.Context, id string) error
	UpdateBook(ctx context.Context, book *models.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) CreateBook(ctx context.Context, user *models.Book) error {
	return r.db.Create(user).Error
}

func (r *bookRepository) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	var book models.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	var books []*models.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) DeleteBook(ctx context.Context, id string) error {
	var book models.Book
	if err := r.db.Delete(&book, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookRepository) UpdateBook(ctx context.Context, book *models.Book) error {
	return r.db.Save(book).Error
}
