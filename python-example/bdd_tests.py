import unittest
from username_validator import UsernameValidator

class BDDTests(unittest.TestCase):
    def setUp(self):
        self.validator = UsernameValidator()
    
    def test_should_accept_valid_usernames(self):
        # Examples of valid usernames
        self.assertTrue(self.validator.is_valid("abc"))
        self.assertTrue(self.validator.is_valid("user123"))
        self.assertTrue(self.validator.is_valid("valid_username"))
    
    def test_should_reject_too_short_usernames(self):
        # Examples of too short usernames
        self.assertFalse(self.validator.is_valid(""))
        self.assertFalse(self.validator.is_valid("ab"))
    
    def test_should_reject_too_long_usernames(self):
        # Examples of too long usernames
        self.assertFalse(self.validator.is_valid("abcdefghijklmnopqrstuvwxyz"))
    
    def test_should_reject_usernames_with_illegal_chars(self):
        # Examples of usernames with illegal chars
        self.assertFalse(self.validator.is_valid("user@name"))
        self.assertFalse(self.validator.is_valid("special$chars"))

if __name__ == "__main__":
    unittest.main()