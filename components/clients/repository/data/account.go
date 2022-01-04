package data

import (
	"database/sql"

	"github.com/crdev13/moneyprocessing/components/clients/entity"
)

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

func ScanAccountFromDBRowResult(
	result *sql.Row,
) (*Account, *string) {
	account := &Account{}
	err := result.Scan(
		&account.ID,
		&account.ClientID,
		&account.Currency,
		&account.Amount,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		message := "Error, cannot find account"
		return nil, &message
	}
	return account, nil
}

func ScanAccountFromDBRowsResult(
	result *sql.Rows,
) (*Account, *string) {
	account := &Account{}
	err := result.Scan(
		&account.ID,
		&account.ClientID,
		&account.Currency,
		&account.Amount,
	)
	if err != nil {
		message := "Error, cannot find account"
		return nil, &message
	}
	return account, nil
}
