package output

import "github.com/crdev13/moneyprocessing/components/clients/entity"

type Account struct {
	ID       uint32  `json:"id"`
	Client   *string `json:"client,omitempty"`
	Currency string  `json:"currency"`
	Amount   float32 `json:"amount"`
}

func MakeAccountOutputFromEntity(data *entity.Account, client *string) *Account {
	if data == nil {
		return nil
	}
	account := &Account{
		ID:       data.ID,
		Client:   client,
		Currency: data.Currency,
		Amount:   data.Amount,
	}
	return account
}
