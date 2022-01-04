package input

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateAccountRequest struct {
	ClientID uint32   `json:"client_id"`
	Account  *Account `json:"account,omitempty"`
}

func (data *CreateAccountRequest) Validate() error {
	if data == nil {
		return fmt.Errorf("Error, no input data")
	}
	if data.ClientID == 0 {
		return fmt.Errorf("Error, invalid client identification")
	}
	if data.Account == nil {
		return fmt.Errorf("Error, no input data in account field")
	}
	if err := validateAccounts([]*Account{data.Account}); err != nil {
		return err
	}
	return nil
}

func MakeCreateAccountRequest(r *http.Request) (*CreateAccountRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var requestForm CreateAccountRequest
	if err := decoder.Decode(&requestForm); err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return &requestForm, nil
}
