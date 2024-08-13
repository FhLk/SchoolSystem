package services

import (
	"Library-backend/models"
	"Library-backend/repositories"
	"context"
)

type AccountService interface {
	CreateAccount(ctx context.Context, account *models.Account) error
	GetAccountByID(ctx context.Context, id string) (*models.Account, error)
	GetAllAccounts(ctx context.Context) ([]*models.Account, error)
	DeleteAccount(ctx context.Context, id string) error
	UpdateAccount(ctx context.Context, account *models.Account) error
}

type AccountServiceImpl struct {
	accountService repositories.AccountRepository
}

func NewAccountService(accountRepository repositories.AccountRepository) *AccountServiceImpl {
	return &AccountServiceImpl{accountService: accountRepository}
}

func (s *AccountServiceImpl) CreateAccount(ctx context.Context, account *models.Account) error {
	return s.accountService.CreateAccount(ctx, account)
}

func (s *AccountServiceImpl) GetAccountByID(ctx context.Context, id string) (*models.Account, error) {
	return s.accountService.GetAccountByID(ctx, id)
}

func (s *AccountServiceImpl) GetAllAccounts(ctx context.Context) ([]*models.Account, error) {
	return s.accountService.GetAllAccount(ctx)
}

func (s *AccountServiceImpl) DeleteAccount(ctx context.Context, id string) error {
	return s.accountService.DeleteAccount(ctx, id)
}

func (s *AccountServiceImpl) UpdateAccount(ctx context.Context, account *models.Account) error {
	return s.accountService.UpdateAccount(ctx, account)
}
