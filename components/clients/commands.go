package clients

import "github.com/crdev13/moneyprocessing/components/clients/dto/output"

type CreateClient interface {
	Execute() (*output.CreateClientResponse, error)
}
