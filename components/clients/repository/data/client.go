package data

import "github.com/crdev13/moneyprocessing/components/clients/entity"

type Client struct {
	ID       uint32
	Name     string
	Accounts []*Account
}

type Account struct {
	ID       uint32
	ClientID uint32
	Currency string
	Amount   float32
}

func (data *Account) ConvertAccountRowResultToEntity() *entity.Account {
	if data == nil {
		return nil
	}
	account := &entity.Account{
		ID:       data.ID,
		ClientID: data.ClientID,
		Currency: data.Currency,
		Amount:   data.Amount,
	}

	return account
}

func (data *Client) ConvertClientRowResultToEntity() *entity.Client {
	if data == nil {
		return nil
	}
	client := &entity.Client{
		ID:   data.ID,
		Name: data.Name,
	}
	accounts := []*entity.Account{}
	for _, acc := range data.Accounts {
		account := acc.ConvertAccountRowResultToEntity()
		accounts = append(accounts, account)
	}
	client.Accounts = accounts
	return client
}
