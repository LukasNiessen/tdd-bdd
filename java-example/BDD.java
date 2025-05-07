package com.example.username;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.junit.jupiter.api.Assertions.assertFalse;

public class BDDTest {
    
    private UsernameValidator validator;
    
    @BeforeEach
    void setUp() {
        validator = new UsernameValidator();
    }
    
    @Test
    void shouldAcceptValidUsernames() {
        // Examples of valid usernames
        assertTrue(validator.isValid("abc"));
        assertTrue(validator.isValid("user123"));
        assertTrue(validator.isValid("valid_username"));
    }

    @Test
    void shouldRejectTooShortUsernames() {
        // Examples of too short usernames
        assertFalse(validator.isValid(""));
        assertFalse(validator.isValid("ab"));
    }

    @Test
    void shouldRejectTooLongUsernames() {
        // Examples of too long usernames
        assertFalse(validator.isValid("abcdefghijklmnopqrstuvwxyz"));
    }

    @Test
    void shouldRejectUsernamesWithIllegalChars() {
        // Examples of usernames with illegal chars
        assertFalse(validator.isValid("user@name"));
        assertFalse(validator.isValid("special$chars"));
    }
}
