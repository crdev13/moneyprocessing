package repository

import (
	"database/sql"
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/entity"
	"github.com/crdev13/moneyprocessing/components/clients/repository/dto/input"
	"github.com/crdev13/moneyprocessing/components/clients/repository/memory"
	"github.com/crdev13/moneyprocessing/components/clients/repository/postgres"
)

type ClientsRepository interface {
	Close()
	HasConnection() bool
	CreateClient(request *input.CreateClient) error
	FindClientByID(clientID uint32) (*entity.Client, error)
	CreateAccount(request *input.CreateAccount) error
	FindAccountByID(accountID uint32) (*entity.Account, error)
}

func NewInMemoryClientsRepository() ClientsRepository {
	clientsRepository := &memory.ClientsRepository{
		Clients:  memory.Clients,
		Accounts: memory.Accounts,
	}
	clientsRepository.InitSequenceID()
	return clientsRepository
}

func NewInPostgreSQLClientsRepository(dbConn *sql.DB) (ClientsRepository, error) {
	if dbConn == nil {
		return nil, fmt.Errorf("Error, no database connection")
	}
	return &postgres.ClientsRepository{DB: dbConn}, nil
}
