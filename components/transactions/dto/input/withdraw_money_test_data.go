package input

type WithdrawMoneyTestCase struct {
	Name      string
	Request   *WithdrawRequest
	NeedError bool
	MsgError  string
}

var WithdrawMoneyTestCases = []*WithdrawMoneyTestCase{
	{
		Name:      "NotWithdrawWithNullInput",
		Request:   nil,
		NeedError: true,
		MsgError:  "Error, no input data",
	},
	{
		Name: "NotWithdrawWithInvalidAccountIdentification",
		Request: &WithdrawRequest{
			AccountID: 0,
			Amount:    10,
		},
		NeedError: true,
		MsgError:  "Error, invalid account identification",
	},
	{
		Name: "NotWithdrawWithInvalidAmount",
		Request: &WithdrawRequest{
			AccountID: 1,
			Amount:    0,
		},
		NeedError: true,
		MsgError:  "Error, invalid amount",
	},
	{
		Name: "NotWithdrawWithInvalidAccount",
		Request: &WithdrawRequest{
			AccountID: 99,
			Amount:    10,
		},
		NeedError: true,
		MsgError:  "Error, you cannot make a withdraw for an account that doesn't exist",
	},
	{
		Name: "NotWithdrawWithNoFunds",
		Request: &WithdrawRequest{
			AccountID: 1,
			Amount:    1000,
		},
		NeedError: true,
		MsgError:  "Error, account has insufficient funds",
	},
	{
		Name: "WithdrawMoney",
		Request: &WithdrawRequest{
			AccountID: 1,
			Amount:    100,
		},
		NeedError: false,
	},
}
