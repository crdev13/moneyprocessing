package output

import "github.com/crdev13/moneyprocessing/components/transactions/entity"

type Transaction struct {
	ID         uint32  `json:"transaction_id"`
	SenderID   *uint32 `json:"sender_account_id,omitempty"`
	ReceiverID *uint32 `json:"reciever_account_id,omitempty"`
	Type       string  `json:"type"`
	Amount     float32 `json:"amount"`
	CreatedAt  string  `json:"created_at"`
}

func MakeTransactionOutputFromEntity(data *entity.Transaction) *Transaction {
	if data == nil {
		return nil
	}
	transaction := &Transaction{
		ID:         data.ID,
		SenderID:   data.SenderID,
		ReceiverID: data.ReceiverID,
		Type:       data.Type,
		Amount:     data.Amount,
		CreatedAt:  data.CreatedAt,
	}
	return transaction
}

func MakeTransactionstOutputFromEntity(data []*entity.Transaction) []*Transaction {
	if len(data) == 0 {
		return nil
	}
	transactions := []*Transaction{}
	for _, tx := range data {
		transaction := MakeTransactionOutputFromEntity(tx)
		transactions = append(transactions, transaction)
	}
	return transactions
}
