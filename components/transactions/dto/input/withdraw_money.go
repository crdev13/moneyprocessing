package input

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WithdrawRequest struct {
	AccountID uint32  `json:"account_id"`
	Amount    float32 `json:"amount"`
}

func (data *WithdrawRequest) Validate() error {
	if data == nil {
		return fmt.Errorf("Error, no input data")
	}
	if data.AccountID == 0 {
		return fmt.Errorf("Error, invalid account identification")
	}
	if err := validateCurrency(data.Amount); err != nil {
		return err
	}
	return nil
}

func MakeWithdrawRequest(r *http.Request) (*WithdrawRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var requestForm WithdrawRequest
	if err := decoder.Decode(&requestForm); err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return &requestForm, nil
}
