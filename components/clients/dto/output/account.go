package output

import (
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/entity"
)

type Account struct {
	ID           uint32      `json:"id"`
	Client       *string     `json:"client,omitempty"`
	Currency     string      `json:"currency"`
	Amount       string      `json:"amount"`
	Transactions interface{} `json:"transactions,omitempty"`
}

func MakeAccountOutputFromEntity(data *entity.Account, client *string) *Account {
	if data == nil {
		return nil
	}
	account := &Account{
		ID:       data.ID,
		Client:   client,
		Currency: data.Currency,
		Amount:   fmt.Sprintf("%0.2f", data.Amount),
	}
	return account
}
