package input

type CreateAccountTestCase struct {
	Name      string
	Request   *CreateAccountRequest
	NeedError bool
	MsgError  string
}

var CreateAccountTestCases = []*CreateAccountTestCase{
	{
		Name:      "NotCreateClientWithNullInput",
		Request:   nil,
		NeedError: true,
		MsgError:  "Error, no input data",
	},
	{
		Name: "NotCreateClientWithNullAccountField",
		Request: &CreateAccountRequest{
			ClientID: 1,
			Account:  nil,
		},
		NeedError: true,
		MsgError:  "Error, no input data in account field",
	},
	{
		Name: "NotCreateClientWithWrongCurrency",
		Request: &CreateAccountRequest{
			ClientID: 1,
			Account: &Account{
				Currency: "BO",
			},
		},
		NeedError: true,
		MsgError:  "Error, you cannot create accounts with invalid currency",
	},
	{
		Name: "NotCreateClientWithInvalidClientIdentification",
		Request: &CreateAccountRequest{
			ClientID: 0,
			Account: &Account{
				Currency: "USD",
			},
		},
		NeedError: true,
		MsgError:  "Error, invalid client identification",
	},
	{
		Name: "NotCreateClientWithNotStoredClient",
		Request: &CreateAccountRequest{
			ClientID: 99,
			Account: &Account{
				Currency: "USD",
			},
		},
		NeedError: true,
		MsgError:  "Error, you cannot create accounts for a client that doesn't exist",
	},
	{
		Name: "NotCreateAccountWithDuplicatedAccount",
		Request: &CreateAccountRequest{
			ClientID: 1,
			Account: &Account{
				Currency: "USD",
			},
		},
		NeedError: true,
		MsgError:  "Error, you cannot create account (a client cannot have duplicated accounts)",
	},
	{
		Name: "NotCreateAccountForClientsWithThreeAccounts",
		Request: &CreateAccountRequest{
			ClientID: 3,
			Account: &Account{
				Currency: "USD",
			},
		},
		NeedError: true,
		MsgError:  "Error, you cannot create accounts for a client that has 3 accounts",
	},
	{
		Name: "CreateAccount",
		Request: &CreateAccountRequest{
			ClientID: 4,
			Account: &Account{
				Currency: "USD",
			},
		},
		NeedError: false,
	},
}
