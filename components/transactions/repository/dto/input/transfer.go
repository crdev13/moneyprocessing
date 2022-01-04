package input

import inputrequest "github.com/crdev13/moneyprocessing/components/transactions/dto/input"

type Transfer struct {
	Sender   *uint32
	Receiver *uint32
	Type     string
	Amount   float32
}

func MakeTransferInputFromRequest(
	request *inputrequest.TransferRequest,
) *Transfer {
	transfer := &Transfer{
		Sender:   &request.SenderID,
		Receiver: &request.ReceiverID,
		Type:     "TRANSFER",
		Amount:   request.Amount,
	}
	return transfer
}
