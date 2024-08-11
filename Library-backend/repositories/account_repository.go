package repositories

import (
	"Library-backend/models"
	"context"
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, Account *models.Account) error
	GetAccountByID(ctx context.Context, id string) (*models.Account, error)
	GetAllAccount(ctx context.Context) ([]*models.Account, error)
	DeleteAccount(ctx context.Context, id string) error
	UpdateAccount(ctx context.Context, Account *models.Account) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) CreateAccount(ctx context.Context, Account *models.Account) error {
	return r.db.Create(Account).Error
}

func (r *accountRepository) GetAccountByID(ctx context.Context, id string) (*models.Account, error) {
	var account models.Account
	if err := r.db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) GetAllAccount(ctx context.Context) ([]*models.Account, error) {
	var accounts []*models.Account
	if err := r.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *accountRepository) DeleteAccount(ctx context.Context, id string) error {
	var account models.Account
	if err := r.db.Delete(&account, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *accountRepository) UpdateAccount(ctx context.Context, Account *models.Account) error {
	return r.db.Save(Account).Error
}
