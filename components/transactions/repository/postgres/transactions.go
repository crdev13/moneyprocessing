package postgres

import (
	"database/sql"
	"fmt"

	accountrepository "github.com/crdev13/moneyprocessing/components/clients/repository/postgres"
	"github.com/crdev13/moneyprocessing/components/transactions/entity"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/data"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
)

func (repository *TransactionsRepository) DepositMoney(request *input.Deposit) error {
	transaction, err := repository.DB.Begin()
	if err != nil {
		return fmt.Errorf("Error, deposit transaction cannot be started")
	}
	if _, err := createTransaction(
		transaction,
		nil,
		request.Receiver,
		request.Type,
		request.Amount,
	); err != nil {
		_ = transaction.Rollback()
		return err
	}
	if err := accountrepository.DepositTx(transaction, request.Receiver, request.Amount); err != nil {
		_ = transaction.Rollback()
		return err
	}
	if err := transaction.Commit(); err != nil {
		return fmt.Errorf("Error, deposit transaction cannot be finished")
	}
	return nil
}
func (repository *TransactionsRepository) WithdrawMoney(request *input.Withdraw) error {
	transaction, err := repository.DB.Begin()
	if err != nil {
		return fmt.Errorf("Error, withdraw transaction cannot be started")
	}
	if _, err := createTransaction(
		transaction,
		request.Sender,
		nil,
		request.Type,
		request.Amount,
	); err != nil {
		_ = transaction.Rollback()
		return err
	}
	if err := accountrepository.WithdrawTx(transaction, request.Sender, request.Amount); err != nil {
		_ = transaction.Rollback()
		return err
	}
	if err := transaction.Commit(); err != nil {
		return fmt.Errorf("Error, withdraw transaction cannot be finished")
	}
	return nil
}
func (repository *TransactionsRepository) TransferMoney(request *input.Transfer) error {
	transaction, err := repository.DB.Begin()
	if err != nil {
		return fmt.Errorf("Error, transfer transaction cannot be started")
	}
	if _, err := createTransaction(
		transaction,
		request.Sender,
		request.Receiver,
		request.Type,
		request.Amount,
	); err != nil {
		_ = transaction.Rollback()
		return err
	}
	if err := accountrepository.TransferTx(transaction, request.Sender, request.Receiver, request.Amount); err != nil {
		_ = transaction.Rollback()
		return err
	}
	if err := transaction.Commit(); err != nil {
		return fmt.Errorf("Error, transfer transaction cannot be finished")
	}
	return nil
}
func (repository *TransactionsRepository) GetTransactionsByAccount(accountID uint32) ([]*entity.Transaction, error) {
	query := `
		SELECT id, sender_account_id, receiver_account_id, type, amount, created_at 
		FROM transactions 
		WHERE(sender_account_id=$1 OR receiver_account_id=$2)
	`
	rows, err := repository.DB.Query(
		query, accountID, accountID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error, cannot find transactions by client(%v)", accountID)
	}
	defer rows.Close()
	transactions := []*entity.Transaction{}
	for rows.Next() {
		data, message := data.ScanTransactionFromDBRowsResult(rows)
		if message != nil {
			return nil, fmt.Errorf("Error, cannot find transactions by account(%v)", accountID)
		}
		transaction := data.ConvertTransactionRowResultToEntity()
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func createTransaction(
	transaction *sql.Tx,
	sender *uint32,
	receiver *uint32,
	typeOfTx string,
	amount float32,
) (uint32, error) {
	query := `
    INSERT INTO transactions (sender_account_id, receiver_account_id, type, amount)
        VALUES ($1, $2, $3, $4)
        RETURNING id;
    `
	row := transaction.QueryRow(
		query, sender, receiver, typeOfTx, amount,
	)
	var transactionID uint32
	err := row.Scan(&transactionID)
	if err != nil {
		fmt.Println(err)
		return transactionID, fmt.Errorf("Error, cannot save transaction")
	}
	return transactionID, nil
}
