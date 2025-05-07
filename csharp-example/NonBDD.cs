using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Moq;
using UsernameValidation;

namespace UsernameValidationTests
{
    [TestClass]
    public class NonBDDTests
    {
        [TestMethod]
        public void TestIsValid()
        {
            // Create a mock for the validator
            var mockValidator = new Mock<UsernameValidator>() { CallBase = true };
            
            string username = "User@123";
            bool result = mockValidator.Object.IsValid(username);
            
            // Check if all methods were called with the right input
            mockValidator.Verify(v => v.IsTooShort(username), Times.Once);
            mockValidator.Verify(v => v.IsTooLong(username), Times.Once);
            mockValidator.Verify(v => v.ContainsIllegalChars(username), Times.Once);
            
            // Now check if they return the correct values
            Assert.IsFalse(mockValidator.Object.IsTooShort(username));
            Assert.IsFalse(mockValidator.Object.IsTooLong(username));
            Assert.IsTrue(mockValidator.Object.ContainsIllegalChars(username));
        }
    }
}