package input

import (
	inputrequest "github.com/crdev13/moneyprocessing/components/clients/dto/input"
)

type CreateClient struct {
	ClientID uint32
	Name     string
	Accounts []*Account
}

type Account struct {
	ClientID uint32
	Currency string
}

func MakeCreateClientInputFromRequest(
	request *inputrequest.CreateClientRequest,
) *CreateClient {
	client := &CreateClient{
		Name: request.Name,
	}
	accounts := []*Account{}
	for _, acc := range request.Accounts {
		account := &Account{Currency: acc.Currency}
		accounts = append(accounts, account)
	}
	client.Accounts = accounts
	return client
}

func (data *CreateClient) SetClientID(clientID uint32) {
	data.ClientID = clientID
	for _, acc := range data.Accounts {
		acc.ClientID = clientID
	}
}
