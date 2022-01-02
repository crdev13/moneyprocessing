package input

type Transaction struct {
	Sender   *uint32
	Reciever *uint32
	Type     string
	Amount   float32
}
