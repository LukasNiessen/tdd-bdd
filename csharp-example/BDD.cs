using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using UsernameValidation;

namespace UsernameValidationTests
{
    [TestClass]
    public class BDDTests
    {
        private UsernameValidator validator;

        [TestInitialize]
        public void Setup()
        {
            validator = new UsernameValidator();
        }

        [TestMethod]
        public void ShouldAcceptValidUsernames()
        {
            // Examples of valid usernames
            Assert.IsTrue(validator.IsValid("abc"));
            Assert.IsTrue(validator.IsValid("user123"));
            Assert.IsTrue(validator.IsValid("valid_username"));
        }

        [TestMethod]
        public void ShouldRejectTooShortUsernames()
        {
            // Examples of too short usernames
            Assert.IsFalse(validator.IsValid(""));
            Assert.IsFalse(validator.IsValid("ab"));
        }

        [TestMethod]
        public void ShouldRejectTooLongUsernames()
        {
            // Examples of too long usernames
            Assert.IsFalse(validator.IsValid("abcdefghijklmnopqrstuvwxyz"));
        }

        [TestMethod]
        public void ShouldRejectUsernamesWithIllegalChars()
        {
            // Examples of usernames with illegal chars
            Assert.IsFalse(validator.IsValid("user@name"));
            Assert.IsFalse(validator.IsValid("special$chars"));
        }
    }
}