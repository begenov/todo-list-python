package repository

import "database/sql"

type AccountsRepo struct {
	db *sql.DB
}

func NewAccountsRepo(db *sql.DB) *AccountsRepo {
	return &AccountsRepo{
		db: db,
	}
}
