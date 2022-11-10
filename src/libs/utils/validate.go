package utils

import (
	"fmt"
	"net"
	"regexp"

	"strings"

	"github.com/hesahesa/pwdbro"
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

var phoneRegex = regexp.MustCompile(`[+]{1}[1-9]{1}[0-9]{1,13}$`)

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

// IsStrongPassword checks if the password is strong enough
func IsStrongPassword(password string) bool {
	pwdbro := pwdbro.NewDefaultPwdBro()
	status, err := pwdbro.RunParallelChecks(password)
	if err != nil {
		fmt.Println(err)
	}

	for _, resp := range status {
		if !resp.Safe {
			return false
		}
	}

	return true
}
