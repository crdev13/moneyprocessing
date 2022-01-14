package repository

import (
	"database/sql"
	"fmt"

	"github.com/crdev13/moneyprocessing/components/transactions/entity"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/memory"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/postgres"
)

type TransactionsRepository interface {
	Close()
	HasConnection() bool
	DepositMoney(request *input.Deposit) error
	WithdrawMoney(request *input.Withdraw) error
	TransferMoney(request *input.Transfer) error
	GetTransactionsByAccount(accountID uint32) ([]*entity.Transaction, error)
}

func NewInMemoryTransactionsRepository() *memory.TransactionsRepository {
	transactionsRepository := &memory.TransactionsRepository{
		Transactions: memory.Transactions,
	}
	transactionsRepository.InitSequenceID()
	return transactionsRepository
}

func NewInPostgreSQLTransactionsRepository(dbConn *sql.DB) (*postgres.TransactionsRepository, error) {
	if dbConn == nil {
		return nil, fmt.Errorf("Error, no database connection")
	}
	return &postgres.TransactionsRepository{DB: dbConn}, nil
}
