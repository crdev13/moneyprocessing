package clients

import (
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/dto/input"
	"github.com/crdev13/moneyprocessing/components/clients/dto/output"
	"github.com/crdev13/moneyprocessing/components/clients/repository"
	inputrepository "github.com/crdev13/moneyprocessing/components/clients/repository/dto/input"
)

type createClient struct {
	clientsRepository repository.ClientsRepository
	request           *input.CreateClientRequest
}

func NewCreateClient(
	clientsRepository repository.ClientsRepository,
	request *input.CreateClientRequest,
) (CreateClient, error) {
	if !clientsRepository.HasConnection() {
		return nil, fmt.Errorf("Error, no database connection")
	}
	if err := request.Validate(); err != nil {
		return nil, err
	}
	return &createClient{
		clientsRepository: clientsRepository,
		request:           request,
	}, nil
}

func (data *createClient) prepareRequestToStoreInRepository() *inputrepository.CreateClient {
	return inputrepository.MakeCreateClientInputFromRequest(data.request)
}

func (data *createClient) Execute() (*output.CreateClientResponse, error) {
	request := data.prepareRequestToStoreInRepository()
	if err := data.clientsRepository.CreateClient(request); err != nil {
		return nil, err
	}
	return &output.CreateClientResponse{
		ClientID: request.ClientID,
	}, nil
}
