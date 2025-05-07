const assert = require("assert");
const sinon = require("sinon");
const UsernameValidator = require("./UsernameValidator");

describe("Username Validator Non-BDD Style", () => {
  it("tests isValid method", () => {
    // Create spy/mock
    const validator = new UsernameValidator();
    const isTooShortSpy = sinon.spy(validator, "isTooShort");
    const isTooLongSpy = sinon.spy(validator, "isTooLong");
    const containsIllegalCharsSpy = sinon.spy(
      validator,
      "containsIllegalChars"
    );

    const username = "User@123";
    const result = validator.isValid(username);

    // Check if all methods were called with the right input
    assert(isTooShortSpy.calledWith(username));
    assert(isTooLongSpy.calledWith(username));
    assert(containsIllegalCharsSpy.calledWith(username));

    // Now check if they return the correct results
    assert.strictEqual(validator.isTooShort(username), false);
    assert.strictEqual(validator.isTooLong(username), false);
    assert.strictEqual(validator.containsIllegalChars(username), true);
  });
});
