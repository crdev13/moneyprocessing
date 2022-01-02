package entity

type Transaction struct {
	ID         uint32
	SenderID   *uint32
	ReceiverID *uint32
	Type       string
	Amount     float32
	CreatedAt  string
}
