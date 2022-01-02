package memory

import (
	"github.com/crdev13/moneyprocessing/components/clients/entity"
)

func (repository *ClientsRepository) FindClientByID(clientID uint32) (*entity.Client, error) {
	clientToConvert, ok := repository.Clients[clientID]
	if !ok {
		return nil, nil
	}
	client := clientToConvert.ConvertClientRowResultToEntity()
	return client, nil
}
