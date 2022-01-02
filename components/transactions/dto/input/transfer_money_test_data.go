package input

type TransferMoneyTestCase struct {
	Name      string
	Request   *TransferRequest
	NeedError bool
	MsgError  string
}

var TransferMoneyTestCases = []*TransferMoneyTestCase{
	{
		Name:      "NotTransferWithNullInput",
		Request:   nil,
		NeedError: true,
		MsgError:  "Error, no input data",
	},
	{
		Name: "NotTransferWithInvalidSenderAccountIdentification",
		Request: &TransferRequest{
			SenderID:   0,
			RecieverID: 2,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, invalid account(sender) identification",
	},
	{
		Name: "NotTransferWithInvalidRecieverAccountIdentification",
		Request: &TransferRequest{
			SenderID:   1,
			RecieverID: 0,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, invalid account(reciever) identification",
	},
	{
		Name: "NotTransferWithTheSameAccountIDs",
		Request: &TransferRequest{
			SenderID:   1,
			RecieverID: 1,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, invalid accounts, they cannot be the same",
	},
	{
		Name: "NotTransferWithInvalidAmount",
		Request: &TransferRequest{
			SenderID:   1,
			RecieverID: 2,
			Amount:     0,
		},
		NeedError: true,
		MsgError:  "Error, invalid amount",
	},
	{
		Name: "NotTransferWithInvalidSenderAccount",
		Request: &TransferRequest{
			SenderID:   99,
			RecieverID: 2,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, you cannot make a transfer for an account that doesn't exist",
	},
	{
		Name: "NotTransferWithInvalidRecieverAccount",
		Request: &TransferRequest{
			SenderID:   2,
			RecieverID: 99,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, you cannot make a transfer for an account that doesn't exist",
	},
	{
		Name: "NotTransferWithNoFunds",
		Request: &TransferRequest{
			SenderID:   1,
			RecieverID: 2,
			Amount:     1000,
		},
		NeedError: true,
		MsgError:  "Error, account has insufficient funds",
	},
	{
		Name: "NotTransferWithDifferentCurrency",
		Request: &TransferRequest{
			SenderID:   3,
			RecieverID: 1,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, you cannot make a transger with acounts that have different currency",
	},
	{
		Name: "TransferMoney",
		Request: &TransferRequest{
			SenderID:   2,
			RecieverID: 1,
			Amount:     10,
		},
		NeedError: false,
	},
}
