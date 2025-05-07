require 'minitest/autorun'
require 'mocha/minitest'
require_relative 'username_validator'

class NonBDDTests < Minitest::Test
  def test_is_valid
    # Create a mock/spy for the validator
    validator = UsernameValidator.new
    
    # Mock the methods to track calls
    validator.expects(:is_too_short).once.returns(false).with('User@123')
    validator.expects(:is_too_long).once.returns(false).with('User@123')
    validator.expects(:contains_illegal_chars).once.returns(true).with('User@123')
    
    username = 'User@123'
    result = validator.is_valid(username)
    
    # The mocked methods will automatically verify that they were called correctly
    # Additionally, we can assert the expected return values
    validator = UsernameValidator.new  # Create a fresh instance for actual return checking
    refute validator.is_too_short(username)
    refute validator.is_too_long(username)
    assert validator.contains_illegal_chars(username)
  end
end