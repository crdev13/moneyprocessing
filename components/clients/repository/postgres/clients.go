package postgres

import (
	"github.com/crdev13/moneyprocessing/components/clients/entity"
	"github.com/crdev13/moneyprocessing/components/clients/repository/dto/input"
)

func (repository *ClientsRepository) CreateClient(request *input.CreateClient) error {
	query := `
    INSERT INTO clients (name)
        VALUES ($1)
        RETURNING id;
    `
	row := repository.DB.QueryRow(
		query, request.Name,
	)
	var clientID uint32
	err := row.Scan(&clientID)
	if err != nil {
		return err
	}
	request.SetClientID(clientID)
	return nil
}
func (repository *ClientsRepository) FindClientByID(clientID uint32) (*entity.Client, error) {
	return nil, nil
}
