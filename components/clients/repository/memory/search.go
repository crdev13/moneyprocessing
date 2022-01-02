package memory

import (
	"fmt"

	"github.com/crdev13/moneyprocessing/components/clients/entity"
)

func (repository *ClientsRepository) FindClientByID(clientID uint32) (*entity.Client, error) {
	clientToConvert, ok := repository.Clients[clientID]
	if !ok {
		return nil, fmt.Errorf("Client not found")
	}
	client := clientToConvert.ConvertClientRowResultToEntity()
	return client, nil
}
