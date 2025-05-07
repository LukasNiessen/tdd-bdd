package com.example.username;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.mockito.Mockito.spy;
import static org.mockito.Mockito.verify;

public class NonBDDTest {
    
    @Test
    public void testIsValid() {
        // create spy / mock
        UsernameValidator validator = spy(new UsernameValidator());

        String username = "User@123";
        boolean result = validator.isValid(username);

        // Check if all methods were called with the right input
        verify(validator).isTooShort(username);
        verify(validator).isTooLong(username);
        verify(validator).containsIllegalChars(username);

        // Now check if they return the correct thing
        assertFalse(validator.isTooShort(username));
        assertFalse(validator.isTooLong(username));
        assertTrue(validator.containsIllegalChars(username));
    }
}