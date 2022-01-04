package clients

import (
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/dto/input"
	"github.com/crdev13/moneyprocessing/components/clients/dto/output"
	"github.com/crdev13/moneyprocessing/components/clients/repository"
	inputrepository "github.com/crdev13/moneyprocessing/components/clients/repository/dto/input"
)

type createAccount struct {
	clientsRepository repository.ClientsRepository
	request           *input.CreateAccountRequest
}

func NewCreateAccount(
	clientsRepository repository.ClientsRepository,
	request *input.CreateAccountRequest,
) (CreateAccount, error) {
	if !clientsRepository.HasConnection() {
		return nil, fmt.Errorf("Error, no database connection")
	}
	if err := request.Validate(); err != nil {
		return nil, err
	}
	return &createAccount{
		clientsRepository: clientsRepository,
		request:           request,
	}, nil
}

func (data *createAccount) prepareRequestToStoreInRepository() *inputrepository.CreateAccount {
	return inputrepository.MakeCreateAccountInputFromRequest(data.request)
}

func (data *createAccount) validate() error {
	client, err := data.clientsRepository.FindClientByID(data.request.ClientID)
	if err != nil {
		return err
	}
	if err := client.CanPerformAction(); err != nil {
		return fmt.Errorf("Error, you cannot create accounts for a client that doesn't exist")
	}
	if len(client.Accounts) == 3 {
		return fmt.Errorf("Error, you cannot create accounts for a client that has 3 accounts")
	}
	for _, acc := range client.Accounts {
		if acc.Currency == data.request.Account.Currency {
			return fmt.Errorf("Error, you cannot create account (a client cannot have duplicated accounts)")
		}
	}
	return nil
}

func (data *createAccount) Execute() (*output.CreateAccountResponse, error) {
	if err := data.validate(); err != nil {
		return nil, err
	}
	request := data.prepareRequestToStoreInRepository()
	if err := data.clientsRepository.CreateAccount(request); err != nil {
		return nil, err
	}
	return &output.CreateAccountResponse{
		ClientID:  request.ClientID,
		AccountID: request.AccountID,
	}, nil
}
