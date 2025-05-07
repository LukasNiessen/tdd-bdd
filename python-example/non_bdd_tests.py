import unittest
from unittest.mock import patch, MagicMock
from username_validator import UsernameValidator

class NonBDDTests(unittest.TestCase):
    def test_is_valid(self):
        # Create a mock/spy for the validator
        validator = UsernameValidator()
        
        # Mock the methods to track calls
        validator.is_too_short = MagicMock(wraps=validator.is_too_short)
        validator.is_too_long = MagicMock(wraps=validator.is_too_long)
        validator.contains_illegal_chars = MagicMock(wraps=validator.contains_illegal_chars)
        
        username = "User@123"
        result = validator.is_valid(username)
        
        # Check if all methods were called with the right input
        validator.is_too_short.assert_called_once_with(username)
        validator.is_too_long.assert_called_once_with(username)
        validator.contains_illegal_chars.assert_called_once_with(username)
        
        # Now check if they return the correct results
        self.assertFalse(validator.is_too_short(username))
        self.assertFalse(validator.is_too_long(username))
        self.assertTrue(validator.contains_illegal_chars(username))

if __name__ == "__main__":
    unittest.main()