require 'minitest/autorun'
require_relative 'username_validator'

class BDDTests < Minitest::Test
  def setup
    @validator = UsernameValidator.new
  end
  
  def test_should_accept_valid_usernames
    # Examples of valid usernames
    assert @validator.is_valid('abc')
    assert @validator.is_valid('user123')
    assert @validator.is_valid('valid_username')
  end
  
  def test_should_reject_too_short_usernames
    # Examples of too short usernames
    refute @validator.is_valid('')
    refute @validator.is_valid('ab')
  end
  
  def test_should_reject_too_long_usernames
    # Examples of too long usernames
    refute @validator.is_valid('abcdefghijklmnopqrstuvwxyz')
  end
  
  def test_should_reject_usernames_with_illegal_chars
    # Examples of usernames with illegal chars
    refute @validator.is_valid('user@name')
    refute @validator.is_valid('special$chars')
  end
end