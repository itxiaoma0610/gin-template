package serializer

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TokenData struct {
	User      interface{} `json:"user"`
	Token     string      `json:"token"`
	TokenType string      `json:"token_type"`
}
