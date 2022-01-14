package entity

import "github.com/shopspring/decimal"

type Transaction struct {
	ID         uint32
	SenderID   *uint32
	ReceiverID *uint32
	Type       string
	Amount     decimal.Decimal
	CreatedAt  string
}
