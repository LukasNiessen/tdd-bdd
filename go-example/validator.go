package validator

import (
	"regexp"
)

// UsernameValidator provides methods to validate usernames
type UsernameValidator struct{}

// IsValid checks if a username is valid according to all rules
func (v *UsernameValidator) IsValid(username string) bool {
	if v.IsTooShort(username) {
		return false
	}
	if v.IsTooLong(username) {
		return false
	}
	if v.ContainsIllegalChars(username) {
		return false
	}
	return true
}

// IsTooShort checks if username is too short (less than 3 characters)
func (v *UsernameValidator) IsTooShort(username string) bool {
	return len(username) < 3
}

// IsTooLong checks if username is too long (more than 20 characters)
func (v *UsernameValidator) IsTooLong(username string) bool {
	return len(username) > 20
}

// ContainsIllegalChars checks if username contains illegal characters
// only alphanumeric and underscores are allowed
func (v *UsernameValidator) ContainsIllegalChars(username string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", username)
	return !matched
}