package memory

import (
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/repository/data"
	"github.com/crdev13/moneyprocessing/components/clients/repository/dto/input"
)

func (repository *ClientsRepository) CreateAccount(request *input.CreateAccount) error {
	repository.nextSequenceAccountID()
	accID := repository.AccountsSequenceID
	account := &data.Account{
		ID:       accID,
		ClientID: request.ClientID,
		Currency: request.Currency,
		Amount:   0,
	}
	client, ok := repository.Clients[request.ClientID]
	if !ok {
		return fmt.Errorf("Error, cannot store new account")
	}
	accounts := []*data.Account{account}
	if len(client.Accounts) != 0 {
		accounts = append(accounts, client.Accounts...)
	}
	client.Accounts = accounts
	repository.Accounts[accID] = account
	request.SetAccountID(accID)
	return nil
}
