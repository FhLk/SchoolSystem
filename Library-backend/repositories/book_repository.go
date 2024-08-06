package repositories

import (
	"Library-backend/models"
	"context"
	"gorm.io/gorm"
)

type BookRepository interface {
	CreateUser(ctx context.Context, user *models.Book) error
	GetUserByID(ctx context.Context, id uint) (*models.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) CreateUser(ctx context.Context, user *models.Book) error {
	return r.db.Create(user).Error
}

func (r *bookRepository) GetUserByID(ctx context.Context, id uint) (*models.Book, error) {
	var user models.Book
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
