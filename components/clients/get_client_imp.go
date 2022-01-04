package clients

import (
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/dto/output"
	"github.com/crdev13/moneyprocessing/components/clients/repository"
)

type getClient struct {
	clientsRepository repository.ClientsRepository
	clientID          uint32
}

func NewGetClientByID(
	clientsRepository repository.ClientsRepository,
	clientID uint32,
) (GetClientByID, error) {
	if !clientsRepository.HasConnection() {
		return nil, fmt.Errorf("Error, no database connection")
	}
	if clientID == 0 {
		return nil, fmt.Errorf("Error, invalid client identification")
	}
	return &getClient{
		clientsRepository: clientsRepository,
		clientID:          clientID,
	}, nil
}

func (data *getClient) Execute() (*output.Client, error) {
	client, err := data.clientsRepository.FindClientByID(data.clientID)
	if err != nil {
		return nil, err
	}
	return output.MakeClientOutputFromEntity(client), nil
}
