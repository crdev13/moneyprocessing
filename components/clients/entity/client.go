package entity

import "fmt"

type Client struct {
	ID       uint32
	Name     string
	Accounts []*Account
}

type Account struct {
	ID       uint32
	ClientID uint32
	Currency string
	Amount   float32
}

func (data *Client) CanPerformAction() error {
	if data == nil {
		return fmt.Errorf("Error, client cannot perform action")
	}
	return nil
}

func (data *Account) CanPerformAction() error {
	if data == nil {
		return fmt.Errorf("Error, account cannot perform action")
	}
	return nil
}
