package validator_test

import (
	"testing"
	
	"github.com/example/validator"
	"github.com/stretchr/testify/mock"
)

// MockValidator is a mock implementation of UsernameValidator for testing
type MockValidator struct {
	mock.Mock
	validator.UsernameValidator
}

// IsTooShort mocks the IsTooShort method
func (m *MockValidator) IsTooShort(username string) bool {
	args := m.Called(username)
	return args.Bool(0)
}

// IsTooLong mocks the IsTooLong method
func (m *MockValidator) IsTooLong(username string) bool {
	args := m.Called(username)
	return args.Bool(0)
}

// ContainsIllegalChars mocks the ContainsIllegalChars method
func (m *MockValidator) ContainsIllegalChars(username string) bool {
	args := m.Called(username)
	return args.Bool(0)
}

func TestUsernameValidatorNonBDD(t *testing.T) {
	t.Run("test IsValid implementation details", func(t *testing.T) {
		// Create a mock validator
		mockValidator := new(MockValidator)
		username := "User@123"
		
		// Set expectations
		mockValidator.On("IsTooShort", username).Return(false)
		mockValidator.On("IsTooLong", username).Return(false)
		mockValidator.On("ContainsIllegalChars", username).Return(true)
		
		// Call IsValid
		result := mockValidator.IsValid(username)
		
		// Verify expectations
		mockValidator.AssertExpectations(t)
		
		// Test with a real validator to verify the expected outcomes
		realValidator := &validator.UsernameValidator{}
		if realValidator.IsTooShort(username) {
			t.Error("Expected username not to be too short")
		}
		if realValidator.IsTooLong(username) {
			t.Error("Expected username not to be too long")
		}
		if !realValidator.ContainsIllegalChars(username) {
			t.Error("Expected username to contain illegal characters")
		}
	})
}