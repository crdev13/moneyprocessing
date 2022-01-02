package posgres

import (
	"github.com/crdev13/moneyprocessing/components/clients/entity"
	"github.com/crdev13/moneyprocessing/components/clients/repository/dto/input"
)

func (repository *TransactionsRepository) CreateClient(request *input.CreateClient) error {
	return nil
}

func (repository *TransactionsRepository) FindClientByID(clientID uint32) (*entity.Client, error) {
	return nil, nil
}

func (repository *TransactionsRepository) CreateAccount(request *input.CreateAccount) error {
	return nil
}

func (repository *TransactionsRepository) FindAccountByID(accountID uint32) (*entity.Account, error) {
	return nil, nil
}
