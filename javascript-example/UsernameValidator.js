class UsernameValidator {
  isValid(username) {
    if (this.isTooShort(username)) {
      return false;
    }
    if (this.isTooLong(username)) {
      return false;
    }
    if (this.containsIllegalChars(username)) {
      return false;
    }
    return true;
  }

  isTooShort(username) {
    return username.length < 3;
  }

  isTooLong(username) {
    return username.length > 20;
  }

  // allows only alphanumeric and underscores
  containsIllegalChars(username) {
    return !username.match(/^[a-zA-Z0-9_]+$/);
  }
}

module.exports = UsernameValidator;
