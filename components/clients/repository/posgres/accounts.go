package posgres

import (
	"github.com/crdev13/moneyprocessing/components/clients/entity"
	"github.com/crdev13/moneyprocessing/components/clients/repository/dto/input"
)

func (repository *ClientsRepository) CreateAccount(request *input.CreateAccount) error {
	return nil
}
func (repository *ClientsRepository) FindAccountByID(accountID uint32) (*entity.Account, error) {
	return nil, nil
}
