package transactions

type Deposit interface {
	Execute() error
}

type Withdraw interface {
	Execute() error
}

type Transfer interface {
	Execute() error
}
