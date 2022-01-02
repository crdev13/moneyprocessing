package input

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var validCurrencies = map[string]bool{"USD": true, "COP": true, "MXN": true}

type CreateClientRequest struct {
	Name     string     `json:"name"`
	Accounts []*Account `json:"accounts,omitempty"`
}

type Account struct {
	Currency string `json:"currency"`
}

func (data *CreateClientRequest) Validate() error {
	if data == nil {
		return fmt.Errorf("Error, no input data")
	}
	if err := validateAccounts(data.Accounts); err != nil {
		return err
	}
	return nil
}

func validateAccounts(accs []*Account) error {
	numberOfAccounts := len(accs)
	if numberOfAccounts == 0 {
		return nil
	}
	if numberOfAccounts > 3 {
		return fmt.Errorf("Error, you cannot create more than 3 accounts for a client")
	}
	accounts := map[string]bool{}
	hasDuplicated := false
	for idx, acc := range accs {
		currency := strings.ToUpper(acc.Currency)
		_, ok := accounts[currency]
		if ok {
			hasDuplicated = true
			break
		}
		accounts[currency] = true
		accs[idx].Currency = currency
	}
	if hasDuplicated {
		return fmt.Errorf("Error, you cannot create accounts with the same currency")
	}
	notValidCurrency := false
	for currency := range accounts {
		_, ok := validCurrencies[currency]
		if !ok {
			notValidCurrency = true
			break
		}
	}
	if notValidCurrency {
		return fmt.Errorf("Error, you cannot create accounts with invalid currency")
	}
	return nil
}

func MakeCreateClientRequest(r *http.Request) (*CreateClientRequest, error) {
	decoder := json.NewDecoder(r.Body)
	var requestForm CreateClientRequest
	if err := decoder.Decode(&requestForm); err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return &requestForm, nil
}
