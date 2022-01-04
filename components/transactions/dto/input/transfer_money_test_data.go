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
			ReceiverID: 2,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, invalid account(sender) identification",
	},
	{
		Name: "NotTransferWithInvalidReceiverAccountIdentification",
		Request: &TransferRequest{
			SenderID:   1,
			ReceiverID: 0,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, invalid account(receiver) identification",
	},
	{
		Name: "NotTransferWithTheSameAccountIDs",
		Request: &TransferRequest{
			SenderID:   1,
			ReceiverID: 1,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, invalid accounts, they cannot be the same",
	},
	{
		Name: "NotTransferWithInvalidAmount",
		Request: &TransferRequest{
			SenderID:   1,
			ReceiverID: 2,
			Amount:     0,
		},
		NeedError: true,
		MsgError:  "Error, invalid amount",
	},
	{
		Name: "NotTransferWithInvalidSenderAccount",
		Request: &TransferRequest{
			SenderID:   99,
			ReceiverID: 2,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, you cannot make a transfer for an account that doesn't exist",
	},
	{
		Name: "NotTransferWithInvalidReceiverAccount",
		Request: &TransferRequest{
			SenderID:   2,
			ReceiverID: 99,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, you cannot make a transfer for an account that doesn't exist",
	},
	{
		Name: "NotTransferWithNoFunds",
		Request: &TransferRequest{
			SenderID:   1,
			ReceiverID: 2,
			Amount:     1000,
		},
		NeedError: true,
		MsgError:  "Error, account has insufficient funds",
	},
	{
		Name: "NotTransferWithDifferentCurrency",
		Request: &TransferRequest{
			SenderID:   3,
			ReceiverID: 1,
			Amount:     10,
		},
		NeedError: true,
		MsgError:  "Error, you cannot make a transger with acounts that have different currency",
	},
	{
		Name: "TransferMoney",
		Request: &TransferRequest{
			SenderID:   2,
			ReceiverID: 1,
			Amount:     10,
		},
		NeedError: false,
	},
}
