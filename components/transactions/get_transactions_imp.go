package transactions

import (
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients"
	"github.com/crdev13/moneyprocessing/components/clients/dto/output"
	"github.com/crdev13/moneyprocessing/components/clients/repository"
	txsoutput "github.com/crdev13/moneyprocessing/components/transactions/dto/output"
	transactionsrepository "github.com/crdev13/moneyprocessing/components/transactions/repository"
)

type getTransactions struct {
	clientsRepository      repository.ClientsRepository
	transactionsRepository transactionsrepository.TransactionsRepository
	accountID              uint32
}

func NewGetTransactionsByAccount(
	clientsRepository repository.ClientsRepository,
	transactionsRepository transactionsrepository.TransactionsRepository,
	accountID uint32,
) (GetTransactionsByAccount, error) {
	if !clientsRepository.HasConnection() {
		return nil, fmt.Errorf("Error, no database connection")
	}
	if accountID == 0 {
		return nil, fmt.Errorf("Error, invalid account identification")
	}
	return &getTransactions{
		clientsRepository:      clientsRepository,
		transactionsRepository: transactionsRepository,
		accountID:              accountID,
	}, nil
}

func (data *getTransactions) Execute() (*output.Account, error) {
	getAccountCommand, err := clients.NewGetAccountByID(data.clientsRepository, data.accountID)
	if err != nil {
		return nil, err
	}
	account, err := getAccountCommand.Execute()
	if err != nil {
		return nil, err
	}
	txs, err := data.transactionsRepository.GetTransactionsByAccount(account.ID)
	if err != nil {
		return nil, err
	}
	transactions := txsoutput.MakeTransactionstOutputFromEntity(txs)
	account.Transactions = transactions
	return account, nil
}
