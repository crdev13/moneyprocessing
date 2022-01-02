package transactions

type Deposit interface {
	Execute() error
}
