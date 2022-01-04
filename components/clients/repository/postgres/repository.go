package postgres

import (
	"database/sql"
)

type ClientsRepository struct {
	DB *sql.DB
}

func (repository *ClientsRepository) Close() {
	repository.DB.Close()
}

func (repository *ClientsRepository) HasConnection() bool {
	return repository.DB != nil
}
