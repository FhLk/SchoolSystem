package services

import (
	"Library-backend/models"
	"Library-backend/repositories"
	"context"
)

type LibrarianService interface {
	CreateLibrarian(ctx context.Context, librarian *models.Librarian) error
	GetLibrarianByID(ctx context.Context, id string) (*models.Librarian, error)
	GetAllLibrarians(ctx context.Context) ([]*models.Librarian, error)
	DeleteLibrarian(ctx context.Context, id string) error
	UpdateLibrarian(ctx context.Context, librarian *models.Librarian) error
}

type LibrarianServiceImpl struct {
	librarianService repositories.LibrarianRepository
}

func NewLibrarianService(librarianRepository repositories.LibrarianRepository) *LibrarianServiceImpl {
	return &LibrarianServiceImpl{librarianService: librarianRepository}
}

func (s *LibrarianServiceImpl) CreateLibrarian(ctx context.Context, librarian *models.Librarian) error {
	err := s.librarianService.CreateLibrarian(ctx, librarian)
	if err != nil {
		return err
	}
	return nil
}

func (s *LibrarianServiceImpl) GetLibrarianByID(ctx context.Context, id string) (*models.Librarian, error) {
	librarian, err := s.librarianService.GetLibrarianByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return librarian, nil
}

func (s LibrarianServiceImpl) GetAllLibrarians(ctx context.Context) ([]*models.Librarian, error) {
	librarians, err := s.librarianService.GetAllLibrarian(ctx)
	if err != nil {
		return nil, err
	}
	return librarians, nil
}

func (s *LibrarianServiceImpl) DeleteLibrarian(ctx context.Context, id string) error {
	err := s.librarianService.DeleteLibrarian(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *LibrarianServiceImpl) UpdateLibrarian(ctx context.Context, librarian *models.Librarian) error {
	err := s.librarianService.UpdateLibrarian(ctx, librarian)
	if err != nil {
		return err
	}
	return nil
}
