package memory

import (
	clientsrepository "github.com/crdev13/moneyprocessing/components/clients/repository/memory"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/data"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
	"github.com/shopspring/decimal"
)

func (repository *TransactionsRepository) WithdrawMoney(request *input.Withdraw) error {
	repository.nextSequenceTransactionID()
	txID := repository.TransactionsSequenceID
	amount := decimal.NewFromFloat32(request.Amount)
	tx := &data.Transaction{
		ID:       txID,
		SenderID: request.Sender,
		Type:     request.Type,
		Amount:   amount,
	}
	repository.Transactions[txID] = tx
	account, ok := clientsrepository.Accounts[*request.Sender]
	if ok {
		account.Amount = account.Amount.Sub(amount)
	}
	return nil
}
