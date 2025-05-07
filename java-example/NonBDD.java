@Test
public void testIsValidUsername() {
    // create spy / mock
    UsernameValidator validator = spy(new UsernameValidator());

    String username = "User@123";
    boolean result = validator.isValidUsername(username);

    // Check if all methods were called with the right input
    verify(validator).isTooShort(username);
    verify(validator).isTooLong(username);
    verify(validator).containsIllegalCharacters(username);

    // Now check if they return the correct thing
    assertFalse(validator.isTooShort(username));
    assertFalse(validator.isTooLong(username));
    assertTrue(validator.containsIllegalCharacters(username));
}