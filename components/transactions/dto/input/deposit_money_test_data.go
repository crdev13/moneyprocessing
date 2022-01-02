package input

type DepositMoneyTestCase struct {
	Name      string
	Request   *DepositRequest
	NeedError bool
	MsgError  string
}

var DepositMoneyTestCases = []*DepositMoneyTestCase{
	{
		Name:      "NotDepositWithNullInput",
		Request:   nil,
		NeedError: true,
		MsgError:  "Error, no input data",
	},
	{
		Name: "NotDepositWithInvalidAccountIdentification",
		Request: &DepositRequest{
			AccountID: 0,
			Amount:    10,
		},
		NeedError: true,
		MsgError:  "Error, invalid account identification",
	},
	{
		Name: "NotDepositWithInvalidAmount",
		Request: &DepositRequest{
			AccountID: 1,
			Amount:    0,
		},
		NeedError: true,
		MsgError:  "Error, invalid amount",
	},
	{
		Name: "NotDepositWithInvalidAmount",
		Request: &DepositRequest{
			AccountID: 99,
			Amount:    10,
		},
		NeedError: true,
		MsgError:  "Error, you cannot make a deposit for an account that doesn't exist",
	},
	{
		Name: "DepositMoney",
		Request: &DepositRequest{
			AccountID: 1,
			Amount:    100,
		},
		NeedError: false,
	},
}
