@Test
void shouldAcceptValidUsernames() {
    // Examples of valid usernames
    assertTrue(validator.isValidUsername("abc"));
    assertTrue(validator.isValidUsername("user123"));
    ...
}

@Test
void shouldRejectTooShortUsernames() {
    // Examples of too short usernames
    assertFalse(validator.isValidUsername(""));
    assertFalse(validator.isValidUsername("ab"));
    ...
}

@Test
void shouldRejectTooLongUsernames() {
    // Examples of too long usernames
    assertFalse(validator.isValidUsername("abcdefghijklmnopqrstuvwxyz"));
    ...
}

@Test
void shouldRejectUsernamesWithIllegalChars() {
    // Examples of usernames with illegal chars
    assertFalse(validator.isValidUsername("user@name"));
    assertFalse(validator.isValidUsername("special$chars"));
    ...
}
