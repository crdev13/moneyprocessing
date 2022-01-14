package postgres

import (
	"database/sql"
	"sync"
)

type ClientsRepository struct {
	DB *sql.DB
	mu sync.Mutex
}

func (repository *ClientsRepository) Close() {
	repository.DB.Close()
}

func (repository *ClientsRepository) HasConnection() bool {
	return repository.DB != nil
}

func (repository *ClientsRepository) Lock() {
	repository.mu.Lock()
}

func (repository *ClientsRepository) Unlock() {
	repository.mu.Unlock()
}
