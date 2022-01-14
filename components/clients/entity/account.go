package entity

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Account struct {
	ID       uint32
	ClientID uint32
	Currency string
	Amount   decimal.Decimal
}

func (data *Account) CanPerformAction() error {
	if data == nil {
		return fmt.Errorf("Error, account cannot perform action")
	}
	return nil
}

func (data *Account) CanPerformWithdraw(amount float32) error {
	if data.Amount.LessThan(decimal.NewFromFloat32(amount)) {
		return fmt.Errorf("Error, account has insufficient funds")
	}
	return nil
}

func (data *Account) CanRecieveTransfer(sender *Account) error {
	if data.Currency != sender.Currency {
		return fmt.Errorf("Error, you cannot make a transfer with acounts that have different currency")
	}
	return nil
}
