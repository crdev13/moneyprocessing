package memory

import "github.com/crdev13/moneyprocessing/components/transactions/entity"

func (repository *TransactionsRepository) GetTransactionsByAccount(accountID uint32) ([]*entity.Transaction, error) {
	transactions := []*entity.Transaction{}
	for _, tx := range repository.Transactions {
		if tx.SenderID != nil {
			if *tx.SenderID == accountID {
				transaction := tx.ConvertTransactionRowResultToEntity()
				transactions = append(transactions, transaction)
			}
		} else if tx.ReceiverID != nil {
			if *tx.ReceiverID == accountID {
				transaction := tx.ConvertTransactionRowResultToEntity()
				transactions = append(transactions, transaction)
			}
		}
	}
	return transactions, nil
}
