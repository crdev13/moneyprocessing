package output

import "github.com/crdev13/moneyprocessing/components/clients/entity"

type Client struct {
	ID       uint32     `json:"id"`
	Name     string     `json:"name"`
	Accounts []*Account `json:"accounts,omitempty"`
}

func MakeClientOutputFromEntity(data *entity.Client) *Client {
	if data == nil {
		return nil
	}
	client := &Client{
		ID:   data.ID,
		Name: data.Name,
	}
	if len(data.Accounts) == 0 {
		return client
	}
	accounts := []*Account{}
	for _, acc := range data.Accounts {
		account := MakeAccountOutputFromEntity(acc, nil)
		accounts = append(accounts, account)
	}
	client.Accounts = accounts
	return client
}
