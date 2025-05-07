using System;
using System.Text.RegularExpressions;

namespace UsernameValidation
{
    public class UsernameValidator
    {
        public bool IsValid(string username)
        {
            if (IsTooShort(username))
            {
                return false;
            }
            if (IsTooLong(username))
            {
                return false;
            }
            if (ContainsIllegalChars(username))
            {
                return false;
            }
            return true;
        }

        public bool IsTooShort(string username)
        {
            return username.Length < 3;
        }

        public bool IsTooLong(string username)
        {
            return username.Length > 20;
        }

        // allows only alphanumeric and underscores
        public bool ContainsIllegalChars(string username)
        {
            return !Regex.IsMatch(username, @"^[a-zA-Z0-9_]+$");
        }
    }
}