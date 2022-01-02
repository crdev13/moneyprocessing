package entity

type Client struct {
	ID       uint32
	Name     string
	Accounts []*Account
}

type Account struct {
	ID       uint32
	ClientID uint32
	Currency string
}
