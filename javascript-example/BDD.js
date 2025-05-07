const assert = require("assert");
const UsernameValidator = require("./UsernameValidator");

describe("Username Validator BDD Style", () => {
  let validator;

  beforeEach(() => {
    validator = new UsernameValidator();
  });

  it("should accept valid usernames", () => {
    // Examples of valid usernames
    assert.strictEqual(validator.isValid("abc"), true);
    assert.strictEqual(validator.isValid("user123"), true);
    assert.strictEqual(validator.isValid("valid_username"), true);
  });

  it("should reject too short usernames", () => {
    // Examples of too short usernames
    assert.strictEqual(validator.isValid(""), false);
    assert.strictEqual(validator.isValid("ab"), false);
  });

  it("should reject too long usernames", () => {
    // Examples of too long usernames
    assert.strictEqual(validator.isValid("abcdefghijklmnopqrstuvwxyz"), false);
  });

  it("should reject usernames with illegal chars", () => {
    // Examples of usernames with illegal chars
    assert.strictEqual(validator.isValid("user@name"), false);
    assert.strictEqual(validator.isValid("special$chars"), false);
  });
});
