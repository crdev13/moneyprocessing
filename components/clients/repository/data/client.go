package data

import (
	"database/sql"

	"github.com/crdev13/moneyprocessing/components/clients/entity"
)

type Client struct {
	ID       uint32
	Name     string
	Accounts []*Account
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

func ScanClientFromDBRowResult(
	result *sql.Row,
) (*Client, *string) {
	client := &Client{}
	err := result.Scan(
		&client.ID,
		&client.Name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		message := "Error, cannot find client"
		return nil, &message
	}
	return client, nil
}
