package memory

import "github.com/crdev13/moneyprocessing/components/clients/repository/data"

var acc1 = &data.Account{
	ID:       1,
	ClientID: 1,
	Currency: "USD",
}
var client1 = &data.Client{
	ID:   1,
	Name: "Carlos Coronado",
	Accounts: []*data.Account{
		acc1,
	},
}

var acc2 = &data.Account{
	ID:       2,
	ClientID: 2,
	Currency: "USD",
}
var acc3 = &data.Account{
	ID:       3,
	ClientID: 2,
	Currency: "COP",
}
var client2 = &data.Client{
	ID:   2,
	Name: "Carlos Aguilar",
	Accounts: []*data.Account{
		acc2,
		acc3,
	},
}

var acc4 = &data.Account{
	ID:       4,
	ClientID: 3,
	Currency: "USD",
}
var acc5 = &data.Account{
	ID:       5,
	ClientID: 3,
	Currency: "COP",
}
var acc6 = &data.Account{
	ID:       6,
	ClientID: 3,
	Currency: "MXN",
}
var client3 = &data.Client{
	ID:   3,
	Name: "Rodrigo Coronado",
	Accounts: []*data.Account{
		acc4,
		acc5,
		acc6,
	},
}

var client4 = &data.Client{
	ID:   4,
	Name: "Rodrigo Carlos",
}

var Clients = map[uint32]*data.Client{
	client1.ID: client1,
	client2.ID: client2,
	client3.ID: client3,
	client4.ID: client4,
}

var Accounts = map[uint32]*data.Account{
	acc1.ID: acc1,
	acc2.ID: acc2,
	acc3.ID: acc3,
	acc4.ID: acc4,
	acc5.ID: acc5,
	acc6.ID: acc6,
}
