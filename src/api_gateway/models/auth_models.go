package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v3"
)

// StuffLoginModel ...
type StuffLoginModel struct {
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

// Validate StuffLoginModel
func (lm *StuffLoginModel) Validate() error {
	return validation.ValidateStruct(
		lm,
		validation.Field(&lm.Password, validation.Required, validation.Length(8, 30), validation.Match(regexp.MustCompile("[a-z]|[A-Z][0-9]"))),
		validation.Field(&lm.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("[+]{1}[1-9]{1}[0-9]{1,13}$"))),
	)
}
