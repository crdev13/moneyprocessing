package input

import "fmt"

type TransferRequest struct {
	SenderID   uint32  `json:"sender_id"`
	RecieverID uint32  `json:"reciever_id"`
	Amount     float32 `json:"amount"`
}

func (data *TransferRequest) Validate() error {
	if data == nil {
		return fmt.Errorf("Error, no input data")
	}
	if data.SenderID == 0 {
		return fmt.Errorf("Error, invalid account(sender) identification")
	}
	if data.RecieverID == 0 {
		return fmt.Errorf("Error, invalid account(reciever) identification")
	}
	if data.SenderID == data.RecieverID {
		return fmt.Errorf("Error, invalid accounts, they cannot be the same")
	}
	if data.Amount == 0 {
		return fmt.Errorf("Error, invalid amount")
	}
	return nil
}
