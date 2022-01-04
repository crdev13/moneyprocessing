package postgres

import (
	"database/sql"
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/entity"
	"github.com/crdev13/moneyprocessing/components/clients/repository/data"
	"github.com/crdev13/moneyprocessing/components/clients/repository/dto/input"
)

func (repository *ClientsRepository) CreateClient(request *input.CreateClient) error {
	transaction, err := repository.DB.Begin()
	if err != nil {
		return fmt.Errorf("Error, transaction cannot be started")
	}
	if err := saveClientTx(transaction, request); err != nil {
		_ = transaction.Rollback()
		return err
	}
	for _, req := range request.Accounts {
		err := saveAccountTx(transaction, req)
		if err != nil {
			_ = transaction.Rollback()
			return err
		}
	}
	if err := transaction.Commit(); err != nil {
		return fmt.Errorf("Error, transaction cannot be finished")
	}
	return nil
}

func (repository *ClientsRepository) FindClientByID(clientID uint32) (*entity.Client, error) {
	query := "SELECT id, name FROM clients WHERE(id=$1)"
	row := repository.DB.QueryRow(
		query, clientID,
	)
	clientToConvert, message := data.ScanClientFromDBRowResult(row)
	if message != nil {
		return nil, fmt.Errorf("%v with id %v", *message, clientID)
	}
	client := clientToConvert.ConvertClientRowResultToEntity()
	if client == nil {
		return nil, nil
	}
	accounts, err := findAccounts(repository.DB, clientID)
	if err != nil {
		return nil, err
	}
	client.Accounts = accounts
	return client, nil
}

func saveClientTx(
	transaction *sql.Tx,
	request *input.CreateClient,
) error {
	query := `
    INSERT INTO clients (name)
        VALUES ($1)
        RETURNING id;
    `
	row := transaction.QueryRow(
		query, request.Name,
	)
	var clientID uint32
	err := row.Scan(&clientID)
	if err != nil {
		return fmt.Errorf("Error, cannot save client")
	}
	request.SetClientID(clientID)
	return nil
}
