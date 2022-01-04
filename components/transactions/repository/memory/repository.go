package memory

import (
	"github.com/crdev13/moneyprocessing/components/transactions/repository/data"
)

type TransactionsRepository struct {
	Transactions           map[uint32]*data.Transaction
	TransactionsSequenceID uint32
}

func (repository *TransactionsRepository) Close() {
}

func (repository *TransactionsRepository) InitSequenceID() {
	var maxIdentificador uint32 = 0
	for identificator := range repository.Transactions {
		if identificator > maxIdentificador {
			maxIdentificador = identificator
		}
	}
	repository.TransactionsSequenceID = maxIdentificador
}

func (repository *TransactionsRepository) nextSequenceTransactionID() {
	sequence := repository.TransactionsSequenceID
	repository.TransactionsSequenceID = sequence + 1
}

func (repository *TransactionsRepository) HasConnection() bool {
	return true
}
