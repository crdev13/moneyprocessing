package clients

import (
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/dto/output"
	"github.com/crdev13/moneyprocessing/components/clients/repository"
)

type getAccount struct {
	clientsRepository repository.ClientsRepository
	accountID         uint32
}

func NewGetAccountByID(
	clientsRepository repository.ClientsRepository,
	accountID uint32,
) (GetAccountByID, error) {
	if !clientsRepository.HasConnection() {
		return nil, fmt.Errorf("Error, no database connection")
	}
	if accountID == 0 {
		return nil, fmt.Errorf("Error, invalid account identification")
	}
	return &getAccount{
		clientsRepository: clientsRepository,
		accountID:         accountID,
	}, nil
}

func (data *getAccount) Execute() (*output.Account, error) {
	account, err := data.clientsRepository.FindAccountByID(data.accountID)
	if err != nil {
		return nil, err
	}
	client, err := data.clientsRepository.FindClientByID(account.ClientID)
	if err != nil {
		return nil, err
	}
	return output.MakeAccountOutputFromEntity(account, &client.Name), nil
}
