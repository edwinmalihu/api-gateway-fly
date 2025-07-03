package model

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Msg      string `json:"messege"`
	Token    string `json:"token"`
}
