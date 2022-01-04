package input

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
	if data.Amount == 0 {
		return fmt.Errorf("Error, invalid amount")
	}
	return nil
}

func MakeDepositRequest(r *http.Request) (*DepositRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var requestForm DepositRequest
	if err := decoder.Decode(&requestForm); err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return &requestForm, nil
}
