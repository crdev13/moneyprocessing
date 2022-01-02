package input

import (
	inputrequest "github.com/crdev13/moneyprocessing/components/transactions/dto/input"
)

type Deposit struct {
	Amount      float32
	Transaction *Transaction
}

func MakeDepositInputFromRequest(
	request *inputrequest.DepositRequest,
) *Deposit {
	deposit := &Deposit{
		Amount: request.Amount,
		Transaction: &Transaction{
			Reciever: &request.AccountID,
			Type:     "DEPOSIT",
			Amount:   request.Amount,
		},
	}
	return deposit
}
