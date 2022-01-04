package input

import (
	inputrequest "github.com/crdev13/moneyprocessing/components/transactions/dto/input"
)

type Withdraw struct {
	Sender *uint32
	Type   string
	Amount float32
}

func MakeWithdrawInputFromRequest(
	request *inputrequest.WithdrawRequest,
) *Withdraw {
	withdraw := &Withdraw{
		Sender: &request.AccountID,
		Type:   "WITHDRAW",
		Amount: request.Amount,
	}
	return withdraw
}
