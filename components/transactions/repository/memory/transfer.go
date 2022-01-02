package memory

import (
	clientsrepository "github.com/crdev13/moneyprocessing/components/clients/repository/memory"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/data"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
)

func (repository *TransactionsRepository) TransferMoney(request *input.Transfer) error {
	repository.nextSequenceTransactionID()
	txID := repository.TransactionsSequenceID
	tx := &data.Transaction{
		ID:         txID,
		SenderID:   request.Sender,
		ReceiverID: request.Reciever,
		Type:       request.Type,
		Amount:     request.Amount,
	}
	repository.Transactions[txID] = tx

	account, ok := clientsrepository.Accounts[*request.Sender]
	if ok {
		account.Amount -= request.Amount
	}
	account, ok = clientsrepository.Accounts[*request.Reciever]
	if ok {
		account.Amount += request.Amount
	}
	return nil
}
