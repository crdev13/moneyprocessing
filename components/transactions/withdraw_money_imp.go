package transactions

import (
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/repository"
	"github.com/crdev13/moneyprocessing/components/transactions/dto/input"
	transactionsrepository "github.com/crdev13/moneyprocessing/components/transactions/repository"
	inputrepository "github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
)

type withdrawMoney struct {
	clientsRepository      repository.ClientsRepository
	transactionsRepository transactionsrepository.TransactionsRepository
	request                *input.WithdrawRequest
}

func NewWithdrawMoney(
	clientsRepository repository.ClientsRepository,
	transactionsRepository transactionsrepository.TransactionsRepository,
	request *input.WithdrawRequest,
) (Withdraw, error) {
	if !clientsRepository.HasConnection() {
		return nil, fmt.Errorf("Error, no database connection")
	}
	if err := request.Validate(); err != nil {
		return nil, err
	}
	return &withdrawMoney{
		clientsRepository:      clientsRepository,
		transactionsRepository: transactionsRepository,
		request:                request,
	}, nil
}

func (data *withdrawMoney) prepareRequestToStoreInRepository() *inputrepository.Withdraw {
	return inputrepository.MakeWithdrawInputFromRequest(data.request)
}

func (data *withdrawMoney) validate() error {
	account, err := data.clientsRepository.FindAccountByID(data.request.AccountID)
	if err != nil {
		return err
	}
	if err := account.CanPerformAction(); err != nil {
		return fmt.Errorf("Error, you cannot make a withdraw for an account that doesn't exist")
	}
	if err := account.CanPerformWithdraw(data.request.Amount); err != nil {
		return err
	}
	return nil
}

func (data *withdrawMoney) Execute() error {
	data.clientsRepository.Lock()
	defer data.clientsRepository.Unlock()
	if err := data.validate(); err != nil {
		return err
	}
	request := data.prepareRequestToStoreInRepository()
	if err := data.transactionsRepository.WithdrawMoney(request); err != nil {
		return err
	}
	return nil
}
