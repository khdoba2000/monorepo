package utils

type ReqSendCode struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type ReqCheckCode struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
