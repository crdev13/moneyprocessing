package input

import (
	inputrequest "github.com/crdev13/moneyprocessing/components/clients/dto/input"
)

type CreateAccount struct {
	ClientID  uint32
	AccountID uint32
	Currency  string
}

func (data *CreateAccount) SetAccountID(accountID uint32) {
	data.AccountID = accountID
}

func MakeCreateAccountInputFromRequest(
	request *inputrequest.CreateAccountRequest,
) *CreateAccount {
	account := &CreateAccount{
		ClientID: request.ClientID,
		Currency: request.Account.Currency,
	}
	return account
}
