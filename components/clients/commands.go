package clients

import "github.com/crdev13/moneyprocessing/components/clients/dto/output"

type CreateClient interface {
	Execute() (*output.CreateClientResponse, error)
}

type GetClientByID interface {
	Execute() (*output.Client, error)
}

type CreateAccount interface {
	Execute() (*output.CreateAccountResponse, error)
}

type GetAccountByID interface {
	Execute() (*output.Account, error)
}
