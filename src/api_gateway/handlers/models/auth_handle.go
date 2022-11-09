package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

// SendSMSModel ...
type SendSMSModel struct {
	PhoneNumber string `json:"phone_number"`
	Type        string `json:"type"`
}

// StandartErrorModel ...
type StandartErrorModel struct {
	ErrorMessage string `json:"error_message"`
}

// RedisData ...
type RedisData struct {
	Value    string `json:"value"`
	Code     string `json:"code"`
	Verified bool   `json:"verified"`
}

// SuccessMessage ...
type SuccessMessage struct {
	Success bool `json:"success"`
}

// VerifyModel ...
type VerifyModel struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}

// LoginModel ...
type LoginModel struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

// Validate LoginModel
func (lm *LoginModel) Validate() error {
	return validation.ValidateStruct(
		lm, validation.Field(&lm.Password, validation.Required, validation.Length(8, 30), validation.Match(regexp.MustCompile("[a-z]|[A-Z][0-9]"))),
		validation.Field(&lm.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("[+]|[0-9]"))),
	)
}

type LoginResponse struct {
	ID           string `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
