package memory

import (
	clientsrepository "github.com/crdev13/moneyprocessing/components/clients/repository/memory"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/data"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
	"github.com/shopspring/decimal"
)

func (repository *TransactionsRepository) TransferMoney(request *input.Transfer) error {
	repository.nextSequenceTransactionID()
	txID := repository.TransactionsSequenceID
	amount := decimal.NewFromFloat32(request.Amount)
	tx := &data.Transaction{
		ID:         txID,
		SenderID:   request.Sender,
		ReceiverID: request.Receiver,
		Type:       request.Type,
		Amount:     amount,
	}
	repository.Transactions[txID] = tx

	account, ok := clientsrepository.Accounts[*request.Sender]
	if ok {
		account.Amount = account.Amount.Sub(amount)
	}
	account, ok = clientsrepository.Accounts[*request.Receiver]
	if ok {
		account.Amount = account.Amount.Add(amount)
	}
	return nil
}
