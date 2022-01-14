package memory

import (
	clientsrepository "github.com/crdev13/moneyprocessing/components/clients/repository/memory"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/data"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
)

func (repository *TransactionsRepository) DepositMoney(request *input.Deposit) error {
	repository.nextSequenceTransactionID()
	txID := repository.TransactionsSequenceID
	tx := &data.Transaction{
		ID:         txID,
		ReceiverID: request.Receiver,
		Type:       request.Type,
		Amount:     request.Amount,
	}
	repository.Transactions[txID] = tx
	account, ok := clientsrepository.Accounts[*request.Receiver]
	if ok {
		account.Amount = account.Amount.Add(request.Amount)
	}
	return nil
}
