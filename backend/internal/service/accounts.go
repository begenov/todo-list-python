package service

import "github.com/begenov/my-project/backend/internal/repository"

type AccountsService struct {
	repo repository.Accounts
}

func NewAccountsService(repo repository.Accounts) *AccountsService {
	return &AccountsService{
		repo: repo,
	}
}
