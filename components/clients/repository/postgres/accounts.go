package postgres

import (
	"database/sql"
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/entity"
	"github.com/crdev13/moneyprocessing/components/clients/repository/data"
	"github.com/crdev13/moneyprocessing/components/clients/repository/dto/input"
)

func (repository *ClientsRepository) CreateAccount(request *input.CreateAccount) error {
	if err := saveAccount(repository.DB, request); err != nil {
		return err
	}
	return nil
}

func (repository *ClientsRepository) FindAccountByID(accountID uint32) (*entity.Account, error) {
	query := "SELECT id, client_id, currency, amount FROM accounts WHERE(id=$1)"
	row := repository.DB.QueryRow(
		query, accountID,
	)
	accountToConvert, message := data.ScanAccountFromDBRowResult(row)
	if message != nil {
		return nil, fmt.Errorf("%v with id %v", *message, accountID)
	}
	account := accountToConvert.ConvertAccountRowResultToEntity()
	if account == nil {
		return nil, nil
	}
	return account, nil
}

func findAccounts(db *sql.DB, clientID uint32) ([]*entity.Account, error) {
	query := "SELECT id, client_id, currency, amount FROM accounts WHERE(client_id=$1)"
	rows, err := db.Query(
		query, clientID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error, cannot find accounts by client(%v)", clientID)
	}
	defer rows.Close()
	accounts := []*entity.Account{}
	for rows.Next() {
		data, message := data.ScanAccountFromDBRowsResult(rows)
		if message != nil {
			return nil, fmt.Errorf("Error, cannot find accounts by client(%v)", clientID)
		}
		acc := data.ConvertAccountRowResultToEntity()
		accounts = append(accounts, acc)
	}
	return accounts, nil
}

func saveAccount(
	db *sql.DB,
	request *input.CreateAccount,
) error {
	query := `
    INSERT INTO accounts (client_id, currency)
        VALUES ($1,$2)
        RETURNING id;
    `
	row := db.QueryRow(
		query, request.ClientID, request.Currency,
	)
	var accountID uint32
	err := row.Scan(&accountID)
	if err != nil {
		return fmt.Errorf("Error, cannot save account")
	}
	request.SetAccountID(accountID)
	return nil
}

func saveAccountTx(
	transaction *sql.Tx,
	request *input.Account,
) error {
	query := `
    INSERT INTO accounts (client_id, currency)
        VALUES ($1,$2)
        RETURNING id;
    `
	row := transaction.QueryRow(
		query, request.ClientID, request.Currency,
	)
	var accountID uint32
	err := row.Scan(&accountID)
	if err != nil {
		return fmt.Errorf("Error, cannot save account")
	}
	return nil
}

func DepositTx(
	transaction *sql.Tx,
	accountID *uint32,
	amount float32,
) error {
	query := `
	UPDATE accounts SET amount = amount + $1 
	WHERE(id = $2)
    `
	_, err := transaction.Exec(
		query, amount, *accountID,
	)
	if err != nil {
		return fmt.Errorf("Error, cannot deposit money account(%v)", *accountID)
	}
	return nil
}

func WithdrawTx(
	transaction *sql.Tx,
	accountID *uint32,
	amount float32,
) error {
	query := `
	UPDATE accounts SET amount = amount - $1 
	WHERE(id = $2)
    `
	_, err := transaction.Exec(
		query, amount, *accountID,
	)
	if err != nil {
		return fmt.Errorf("Error, cannot withdraw money account(%v)", *accountID)
	}
	return nil
}

func TransferTx(
	transaction *sql.Tx,
	senderID *uint32,
	receiverID *uint32,
	amount float32,
) error {
	if err := WithdrawTx(transaction, senderID, amount); err != nil {
		return fmt.Errorf("Error, cannot transfer money account(%v)", *senderID)
	}
	if err := DepositTx(transaction, receiverID, amount); err != nil {
		return fmt.Errorf("Error, cannot transfer money account(%v)", *receiverID)
	}
	return nil
}
