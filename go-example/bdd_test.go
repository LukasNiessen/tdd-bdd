package validator_test

import (
	"testing"

	"github.com/example/validator"
)

func TestUsernameValidatorBDD(t *testing.T) {
	v := &validator.UsernameValidator{}

	t.Run("should accept valid usernames", func(t *testing.T) {
		// Examples of valid usernames
		if !v.IsValid("abc") {
			t.Error("Expected 'abc' to be valid")
		}
		if !v.IsValid("user123") {
			t.Error("Expected 'user123' to be valid")
		}
		if !v.IsValid("valid_username") {
			t.Error("Expected 'valid_username' to be valid")
		}
	})

	t.Run("should reject too short usernames", func(t *testing.T) {
		// Examples of too short usernames
		if v.IsValid("") {
			t.Error("Expected empty string to be invalid")
		}
		if v.IsValid("ab") {
			t.Error("Expected 'ab' to be invalid")
		}
	})

	t.Run("should reject too long usernames", func(t *testing.T) {
		// Examples of too long usernames
		if v.IsValid("abcdefghijklmnopqrstuvwxyz") {
			t.Error("Expected long string to be invalid")
		}
	})

	t.Run("should reject usernames with illegal chars", func(t *testing.T) {
		// Examples of usernames with illegal chars
		if v.IsValid("user@name") {
			t.Error("Expected 'user@name' to be invalid")
		}
		if v.IsValid("special$chars") {
			t.Error("Expected 'special$chars' to be invalid")
		}
	})
}