package transactions

import (
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/repository"
	"github.com/crdev13/moneyprocessing/components/transactions/dto/input"
	transactionsrepository "github.com/crdev13/moneyprocessing/components/transactions/repository"
	inputrepository "github.com/crdev13/moneyprocessing/components/transactions/repository/dto/input"
)

type transferMoney struct {
	clientsRepository      repository.ClientsRepository
	transactionsRepository transactionsrepository.TransactionsRepository
	request                *input.TransferRequest
}

func NewTransferMoney(
	clientsRepository repository.ClientsRepository,
	transactionsRepository transactionsrepository.TransactionsRepository,
	request *input.TransferRequest,
) (Withdraw, error) {
	if !clientsRepository.HasConnection() {
		return nil, fmt.Errorf("Error, no database connection")
	}
	if err := request.Validate(); err != nil {
		return nil, err
	}
	return &transferMoney{
		clientsRepository:      clientsRepository,
		transactionsRepository: transactionsRepository,
		request:                request,
	}, nil
}

func (data *transferMoney) prepareRequestToStoreInRepository() *inputrepository.Transfer {
	return inputrepository.MakeTransferInputFromRequest(data.request)
}

func (data *transferMoney) validate() error {
	sender, err := data.clientsRepository.FindAccountByID(data.request.SenderID)
	if err != nil {
		return err
	}
	if err := sender.CanPerformAction(); err != nil {
		return fmt.Errorf("Error, you cannot make a transfer for an account that doesn't exist")
	}
	if err := sender.CanPerformWithdraw(data.request.Amount); err != nil {
		return err
	}

	reciever, err := data.clientsRepository.FindAccountByID(data.request.RecieverID)
	if err != nil {
		return err
	}
	if err := reciever.CanPerformAction(); err != nil {
		return fmt.Errorf("Error, you cannot make a transfer for an account that doesn't exist")
	}

	if err := reciever.CanRecieveTransfer(sender); err != nil {
		return err
	}
	return nil
}

func (data *transferMoney) Execute() error {
	if err := data.validate(); err != nil {
		return err
	}
	request := data.prepareRequestToStoreInRepository()
	if err := data.transactionsRepository.TransferMoney(request); err != nil {
		return err
	}
	return nil
}
