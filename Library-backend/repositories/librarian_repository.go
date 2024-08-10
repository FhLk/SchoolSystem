package repositories

import (
	"Library-backend/models"
	"context"
	"gorm.io/gorm"
)

type LibrarianRepository interface {
	CreateLibrarian(ctx context.Context, librarian *models.Librarian) error
	GetLibrarianByID(ctx context.Context, id string) (*models.Librarian, error)
	GetAllLibrarian(ctx context.Context) ([]*models.Librarian, error)
	DeleteLibrarian(ctx context.Context, id string) error
	UpdateLibrarian(ctx context.Context, librarian *models.Librarian) error
}

type librarianRepository struct {
	db *gorm.DB
}

func NewLibrarianRepository(db *gorm.DB) LibrarianRepository {
	return &librarianRepository{db: db}
}

func (r *librarianRepository) CreateLibrarian(ctx context.Context, librarian *models.Librarian) error {
	return r.db.Create(librarian).Error
}

func (r *librarianRepository) GetLibrarianByID(ctx context.Context, id string) (*models.Librarian, error) {
	var librarian models.Librarian
	if err := r.db.First(&librarian, id).Error; err != nil {
		return nil, err
	}
	return &librarian, nil
}

func (r *librarianRepository) GetAllLibrarian(ctx context.Context) ([]*models.Librarian, error) {
	var librarians []*models.Librarian
	if err := r.db.Find(&librarians).Error; err != nil {
		return nil, err
	}
	return librarians, nil
}

func (r *librarianRepository) DeleteLibrarian(ctx context.Context, id string) error {
	var librarian models.Librarian
	if err := r.db.Delete(&librarian, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *librarianRepository) UpdateLibrarian(ctx context.Context, librarian *models.Librarian) error {
	return r.db.Save(librarian).Error
}
