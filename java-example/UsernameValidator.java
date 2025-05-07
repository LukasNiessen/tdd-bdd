public class UsernameValidator {

    public boolean isValid(String username) {
        if (isTooShort(username)) {
            return false;
        }
        if (isTooLong(username)) {
            return false;
        }
        if (containsIllegalChars(username)) {
            return false;
        }
        return true;
    }

    boolean isTooShort(String username) {
        return username.length() < 3;
    }

    boolean isTooLong(String username) {
        return username.length() > 20;
    }

    // allows only alphanumeric and underscores
    boolean containsIllegalChars(String username) {
        return !username.matches("^[a-zA-Z0-9_]+$");
    }
}