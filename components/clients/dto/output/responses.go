package output

type CreateClientResponse struct {
	ClientID uint32 `json:"client_id"`
}

type CreateAccountResponse struct {
	ClientID  uint32 `json:"client_id"`
	AccountID uint32 `json:"account_id"`
}
