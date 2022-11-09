package models

type ReqSendCode struct {
	LoginValue string `json:"login_value"`
}

type ReqCheckCode struct {
	LoginValue string `json:"login_value"`
	Code       string `json:"code"`
}
