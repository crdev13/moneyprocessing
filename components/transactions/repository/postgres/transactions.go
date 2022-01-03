package postgres

import (
	"github.com/crdev13/moneyprocessing/components/transactions/entity"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
)

func (repository *TransactionsRepository) DepositMoney(request *input.Deposit) error {
	return nil
}
func (repository *TransactionsRepository) WithdrawMoney(request *input.Withdraw) error {
	return nil
}
func (repository *TransactionsRepository) TransferMoney(request *input.Transfer) error {
	return nil
}
func (repository *TransactionsRepository) GetTransactionsByAccount(accountID uint32) ([]*entity.Transaction, error) {
	return nil, nil
}
