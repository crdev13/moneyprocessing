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
	if len(request.Accounts) == 0 {
		return nil
	}
	accounts := []*data.Account{}
	for _, acc := range request.Accounts {
		repository.nextSequenceAccountID()
		accID := repository.AccountsSequenceID
		account := &data.Account{
			ID:       accID,
			ClientID: acc.ClientID,
			Currency: acc.Currency,
		}
		accounts = append(accounts, account)
		repository.Accounts[accID] = account
	}
	client.Accounts = accounts
	return nil
}
