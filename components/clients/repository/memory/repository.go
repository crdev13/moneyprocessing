package memory

import (
	"github.com/crdev13/moneyprocessing/components/clients/repository/data"
)

type ClientsRepository struct {
	Clients            map[uint32]*data.Client
	Accounts           map[uint32]*data.Account
	ClientsSequenceID  uint32
	AccountsSequenceID uint32
}

func (repository *ClientsRepository) Close() {
}

func (repository *ClientsRepository) InitSequenceID() {
	var maxIdentificador uint32 = 0
	for identificator := range repository.Clients {
		if identificator > maxIdentificador {
			maxIdentificador = identificator
		}
	}
	repository.ClientsSequenceID = maxIdentificador

	maxIdentificador = 0
	for identificator := range repository.Accounts {
		if identificator > maxIdentificador {
			maxIdentificador = identificator
		}
	}
	repository.AccountsSequenceID = maxIdentificador
}

func (repository *ClientsRepository) nextSequenceClientID() {
	sequence := repository.ClientsSequenceID
	repository.ClientsSequenceID = sequence + 1
}

func (repository *ClientsRepository) nextSequenceAccountID() {
	sequence := repository.AccountsSequenceID
	repository.AccountsSequenceID = sequence + 1
}

func (repository *ClientsRepository) HasConnection() bool {
	return true
}
