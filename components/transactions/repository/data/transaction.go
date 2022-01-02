package data

import "github.com/crdev13/moneyprocessing/components/transactions/entity"

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
