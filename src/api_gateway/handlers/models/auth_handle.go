package models

// SendSMSModel ...
type SendSMSModel struct {
	PhoneNumber string `json:"phone_number"`
	Type        string `json:"type"`
}

// StandartErrorModel
type StandartErrorModel struct {
	ErrorMessage string `json:"error_message"`
}
