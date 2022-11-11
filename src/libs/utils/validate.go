package utils

import (
	"monorepo/src/api_gateway/constants"
	"net"
	"regexp"
	"unicode"

	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// IsEmailValid ...
func IsEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

var phoneRegex = regexp.MustCompile(`(0|\\+62|062|62)[0-9]+$`)

// IsPhoneValid
func IsPhoneValid(p string) bool {
	return phoneRegex.MatchString(p)
}

// Validate for both email ad phone
func ValidatePhoneOrEmail(loginValue string) bool {
	var valid bool

	valid = IsPhoneValid(loginValue)
	if valid {
		return true
	}

	valid = IsEmailValid(loginValue)
	if valid {
		return true
	}

	return false
}

// ValidatePassword check if the password meets requiremnts of
// at least 8 characters, at least one alphabetic, and at least one number
func ValidatePassword(p string) error {
	if len(p) < 8 {
		return constants.ErrPasswordTooShort
	}
	if len(p) > 256 {
		return constants.ErrPasswordTooLong
	}
	hasAnyDigit := false
	hasAnyAlphabetic := false
	for _, c := range p {
		if unicode.IsDigit(c) {
			hasAnyDigit = true
		}

		if unicode.IsLetter(c) {
			hasAnyAlphabetic = true
		}
	}

	if !hasAnyDigit {
		return constants.ErrMustContainDigit
	}

	if !hasAnyAlphabetic {
		return constants.ErrMustContainAlphabetic
	}

	return nil
}
