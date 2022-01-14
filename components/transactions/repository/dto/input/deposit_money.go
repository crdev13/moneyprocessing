package input

import (
	inputrequest "github.com/crdev13/moneyprocessing/components/transactions/dto/input"
	"github.com/shopspring/decimal"
)

type Deposit struct {
	Receiver *uint32
	Type     string
	Amount   decimal.Decimal
}

func MakeDepositInputFromRequest(
	request *inputrequest.DepositRequest,
) *Deposit {
	deposit := &Deposit{
		Receiver: &request.AccountID,
		Type:     "DEPOSIT",
		Amount:   decimal.NewFromFloat32(request.Amount),
	}
	return deposit
}
