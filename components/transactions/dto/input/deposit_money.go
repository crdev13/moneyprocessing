package input

import "fmt"

type DepositRequest struct {
	AccountID uint32  `json:"account_id"`
	Amount    float32 `json:"amount"`
}

func (data *DepositRequest) Validate() error {
	if data == nil {
		return fmt.Errorf("Error, no input data")
	}
	if data.AccountID == 0 {
		return fmt.Errorf("Error, invalid account identification")
	}
	if data.Amount <= 0 {
		return fmt.Errorf("Error, invalid amount")
	}
	return nil
}