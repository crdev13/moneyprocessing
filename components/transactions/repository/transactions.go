package repository

import (
	"github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
	"github.com/crdev13/moneyprocessing/components/transactions/repository/memory"
)

type TransactionsRepository interface {
	Close()
	HasConnection() bool
	DepositMoney(request *input.Deposit) error
	WithdrawMoney(request *input.Withdraw) error
}

func NewInMemoryTransactionsRepository() TransactionsRepository {
	transactionsRepository := &memory.TransactionsRepository{
		Transactions: memory.Transactions,
	}
	transactionsRepository.InitSequenceID()
	return transactionsRepository
}

// func NewInPostgreSQLClientsRepository(dbConn *sql.DB) (ClientsRepository, error) {
// 	if dbConn == nil {
// 		return nil, fmt.Errorf("Error, no database connection")
// 	}
// 	return &posgrerepository.ProductsRepository{DB: dbConn}, nil
// }
