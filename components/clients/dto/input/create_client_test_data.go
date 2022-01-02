package input

type CreateClientTestCase struct {
	Name      string
	Request   *CreateClientRequest
	NeedError bool
	MsgError  string
}

var CreateClientTestCases = []*CreateClientTestCase{
	{
		Name:      "NotCreateClientWithNullInput",
		Request:   nil,
		NeedError: true,
		MsgError:  "Error, no input data",
	},
	{
		Name: "NotCreateClientWithMoreThanThreeAccounts",
		Request: &CreateClientRequest{
			Name: "Test Case",
			Accounts: []*Account{
				{
					Currency: "USD",
				},
				{
					Currency: "COP",
				},
				{
					Currency: "MXN",
				},
				{
					Currency: "BO",
				},
			},
		},
		NeedError: true,
		MsgError:  "Error, you cannot create more than 3 accounts for a client",
	},
	{
		Name: "NotCreateClientWithDuplicatedCurrency",
		Request: &CreateClientRequest{
			Name: "Test Case",
			Accounts: []*Account{
				{
					Currency: "USD",
				},
				{
					Currency: "USD",
				},
			},
		},
		NeedError: true,
		MsgError:  "Error, you cannot create accounts with the same currency",
	},
	{
		Name: "NotCreateClientWithWrongCurrency",
		Request: &CreateClientRequest{
			Name: "Test Case",
			Accounts: []*Account{
				{
					Currency: "BO",
				},
			},
		},
		NeedError: true,
		MsgError:  "Error, you cannot create accounts with invalid currency",
	},
	{
		Name: "CreateClient",
		Request: &CreateClientRequest{
			Name: "Carlos Soto",
		},
		NeedError: false,
	},
	{
		Name: "CreateClientWithAccounts",
		Request: &CreateClientRequest{
			Name: "Rodrigo Soto",
			Accounts: []*Account{
				{
					Currency: "USD",
				},
			},
		},
		NeedError: false,
	},
}
