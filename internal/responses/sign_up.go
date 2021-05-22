package responses

type JWT struct {
	Token     string `json:"access_token"`
	TokenType string `json:"type"`
}