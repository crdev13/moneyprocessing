package transactions

import "github.com/crdev13/moneyprocessing/components/clients/dto/output"

type Deposit interface {
	Execute() error
}

type Withdraw interface {
	Execute() error
}

type Transfer interface {
	Execute() error
}

type GetTransactionsByAccount interface {
	Execute() (*output.Account, error)
}
