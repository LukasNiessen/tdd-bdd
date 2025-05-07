import re

class UsernameValidator:
    def is_valid(self, username):
        if self.is_too_short(username):
            return False
        if self.is_too_long(username):
            return False
        if self.contains_illegal_chars(username):
            return False
        return True
    
    def is_too_short(self, username):
        return len(username) < 3
    
    def is_too_long(self, username):
        return len(username) > 20
    
    # allows only alphanumeric and underscores
    def contains_illegal_chars(self, username):
        return not bool(re.match(r'^[a-zA-Z0-9_]+$', username))