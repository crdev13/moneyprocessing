package memory

import (
	"github.com/crdev13/moneyprocessing/components/clients/repository/data"
	"github.com/crdev13/moneyprocessing/components/clients/repository/dto/input"
)

func (repository *ClientsRepository) CreateClient(request *input.CreateClient) error {
	repository.nextSequenceClientID()
	clientID := repository.ClientsSequenceID
	client := &data.Client{
		ID:   clientID,
		Name: request.Name,
	}
	repository.Clients[clientID] = client
	request.SetClientID(clientID)
	if len(repository.Accounts) == 0 {
		return nil
	}
	for _, acc := range repository.Accounts {
		repository.nextSequenceAccountID()
		accID := repository.AccountsSequenceID
		account := &data.Account{
			ID:       accID,
			ClientID: acc.ClientID,
			Currency: acc.Currency,
		}
		repository.Accounts[accID] = account
	}
	return nil
}
