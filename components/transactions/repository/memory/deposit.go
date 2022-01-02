package memory

import (
	"github.com/crdev13/moneyprocessing/components/transactions/repository/data"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
)

func (repository *TransactionsRepository) DepositMoney(request *input.Deposit) error {

	repository.nextSequenceTransactionID()
	txID := repository.TransactionsSequenceID

	tx := &data.Transaction{
		ID:         txID,
		SenderID:   request.Transaction.Reciever,
		ReceiverID: request.Transaction.Reciever,
		Type:       request.Transaction.Type,
		Amount:     request.Transaction.Amount,
	}
	repository.Transactions[txID] = tx
	return nil
}
