package constants

import "errors"

var ErrPasswordTooShort = errors.New("too short")
var ErrPasswordTooLong = errors.New("too long")
var ErrMustContainDigit = errors.New("must contain at least 1 digit")
var ErrMustContainAlphabetic = errors.New("must contain at least 1 alphabetic")

const ()
