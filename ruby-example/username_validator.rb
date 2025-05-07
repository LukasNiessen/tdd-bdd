class UsernameValidator
  def is_valid(username)
    return false if is_too_short(username)
    return false if is_too_long(username)
    return false if contains_illegal_chars(username)
    true
  end

  def is_too_short(username)
    username.length < 3
  end

  def is_too_long(username)
    username.length > 20
  end

  # allows only alphanumeric and underscores
  def contains_illegal_chars(username)
    !username.match?(/^[a-zA-Z0-9_]+$/)
  end
end