package repository

import "database/sql"

type Accounts interface {
}

type Repository struct {
	Accounts Accounts
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Accounts: NewAccountsRepo(db),
	}
}
