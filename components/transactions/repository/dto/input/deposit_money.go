package input

import (
	inputrequest "github.com/crdev13/moneyprocessing/components/transactions/dto/input"
)

type Deposit struct {
	Reciever *uint32
	Type     string
	Amount   float32
}

func MakeDepositInputFromRequest(
	request *inputrequest.DepositRequest,
) *Deposit {
	deposit := &Deposit{
		Reciever: &request.AccountID,
		Type:     "DEPOSIT",
		Amount:   request.Amount,
	}
	return deposit
}
