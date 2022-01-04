package transactions

import (
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/repository"
	"github.com/crdev13/moneyprocessing/components/transactions/dto/input"
	transactionsrepository "github.com/crdev13/moneyprocessing/components/transactions/repository"
	inputrepository "github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
)

type depositMoney struct {
	clientsRepository      repository.ClientsRepository
	transactionsRepository transactionsrepository.TransactionsRepository
	request                *input.DepositRequest
}

func NewDepositMoney(
	clientsRepository repository.ClientsRepository,
	transactionsRepository transactionsrepository.TransactionsRepository,
	request *input.DepositRequest,
) (Deposit, error) {
	if !clientsRepository.HasConnection() {
		return nil, fmt.Errorf("Error, no database connection")
	}
	if err := request.Validate(); err != nil {
		return nil, err
	}
	return &depositMoney{
		clientsRepository:      clientsRepository,
		transactionsRepository: transactionsRepository,
		request:                request,
	}, nil
}

func (data *depositMoney) prepareRequestToStoreInRepository() *inputrepository.Deposit {
	return inputrepository.MakeDepositInputFromRequest(data.request)
}

func (data *depositMoney) validate() error {
	account, err := data.clientsRepository.FindAccountByID(data.request.AccountID)
	if err != nil {
		return err
	}
	if err := account.CanPerformAction(); err != nil {
		return fmt.Errorf("Error, you cannot make a deposit for an account that doesn't exist")
	}
	return nil
}

func (data *depositMoney) Execute() error {
	if err := data.validate(); err != nil {
		return err
	}
	request := data.prepareRequestToStoreInRepository()
	if err := data.transactionsRepository.DepositMoney(request); err != nil {
		return err
	}
	return nil
}
