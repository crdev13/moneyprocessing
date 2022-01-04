package postgres

import (
	"database/sql"
)

type TransactionsRepository struct {
	DB *sql.DB
}

func (repository *TransactionsRepository) Close() {
	repository.DB.Close()
}

func (repository *TransactionsRepository) HasConnection() bool {
	return repository.DB != nil
}
