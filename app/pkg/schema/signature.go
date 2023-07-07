package schema

type SignatureSchema struct {
	TokenID       string `json:"token_id"`
	Message       string `json:"message"`
	Signature     string `json:"signature"`
	WalletAddress string `json:"wallet_address"`
}
