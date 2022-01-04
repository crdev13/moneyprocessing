package data

import (
	"database/sql"

	"github.com/crdev13/moneyprocessing/components/transactions/entity"
)

type Transaction struct {
	ID         uint32
	SenderID   *uint32
	ReceiverID *uint32
	Type       string
	Amount     float32
	CreatedAt  string
}

func (data *Transaction) ConvertTransactionRowResultToEntity() *entity.Transaction {
	if data == nil {
		return nil
	}
	transaction := &entity.Transaction{
		ID:         data.ID,
		SenderID:   data.SenderID,
		ReceiverID: data.ReceiverID,
		Type:       data.Type,
		Amount:     data.Amount,
		CreatedAt:  data.CreatedAt,
	}

	return transaction
}

func ScanTransactionFromDBRowsResult(
	result *sql.Rows,
) (*Transaction, *string) {
	transaction := &Transaction{}
	err := result.Scan(
		&transaction.ID,
		&transaction.SenderID,
		&transaction.ReceiverID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.CreatedAt,
	)
	if err != nil {
		message := "Error, cannot find transaction"
		return nil, &message
	}
	return transaction, nil
}
