# TDD and BDD Explained

_TDD = Test-Driven Development_  
_BDD = Behavior-Driven Development_

## Behavior-Driven Development

BDD is all about the following mindset: **Do not test code. Test behavior.**

So it's a shift of the testing mindset. This is why in BDD, we also introduced new terms:

- **Test suites** become **specifications**,
- **Test cases** become **scenarios**,
- We don't **test code**, we **verify behavior**.

Let's make this clear by an example.

## Java Example

If you are not familiar with Java, look in the repo files for other languages.

```java
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
```

UsernameValidator checks if a username is valid (3-20 characters, alphanumeric and \_). It returns true if all checks pass, else false.

How to test this? Well, if we test if the code does what it does, it might look like this:

```java
@Test
public void testIsValidUsername() {
    // create spy / mock
    UsernameValidator validator = spy(new UsernameValidator());

    String username = "User@123";
    boolean result = validator.isValidUsername(username);

    // Check if all methods were called with the right input
    verify(validator).isTooShort(username);
    verify(validator).isTooLong(username);
    verify(validator).containsIllegalCharacters(username);

    // Now check if they return the correct thing
    assertFalse(validator.isTooShort(username));
    assertFalse(validator.isTooLong(username));
    assertTrue(validator.containsIllegalCharacters(username));
}
```

This is not great. What if we change the logic inside isValidUsername? Let's say we decide to replace `isTooShort()` and `isTooLong()` by a new method `isLengthAllowed()`?

**The test would break**. Because it almost mirros the implementation. Not good. The test is now tightly coupled to the implementation.

In BDD, we just verify the behavior. So, in this case, we just check if we get the wanted outcome:

```java
@Test
void shouldAcceptValidUsernames() {
    // Examples of valid usernames
    assertTrue(validator.isValidUsername("abc"));
    assertTrue(validator.isValidUsername("user123"));
    ...
}

@Test
void shouldRejectTooShortUsernames() {
    // Examples of too short usernames
    assertFalse(validator.isValidUsername(""));
    assertFalse(validator.isValidUsername("ab"));
    ...
}

@Test
void shouldRejectTooLongUsernames() {
    // Examples of too long usernames
    assertFalse(validator.isValidUsername("abcdefghijklmnopqrstuvwxyz"));
    ...
}

@Test
void shouldRejectUsernamesWithIllegalChars() {
    // Examples of usernames with illegal chars
    assertFalse(validator.isValidUsername("user@name"));
    assertFalse(validator.isValidUsername("special$chars"));
    ...
}
```

Much better. If you change the implementation, the tests will not break. They will work as long as the method works.

Implementation is irrelevant, we only specified our wanted behavior. This is why, in BDD, we don't call it a _test suite_ but we call it a _specification_.

Of course this example is very simplified and doesn't cover all aspects of BDD but it clearly illustrates the core of BDD: **testing code vs verifying behavior**.

## Is it about tools?

Many people think BDD is something written in Gherkin syntax with tools like Cucumber or SpecFlow:

```gherkin
Feature: User login
  Scenario: Successful login
    Given a user with valid credentials
    When the user submits login information
    Then they should be authenticated and redirected to the dashboard
```

While these tools are great and definitely help to implement BDD, it's not limited to them. BDD is much broader. BDD is about behavior, not about tools. You can use BDD with these tools, but also with other tools. Or without tools at all.

## More on BDD

https://www.youtube.com/watch?v=Bq_oz7nCNUA (by Dave Farley)  
https://www.thoughtworks.com/en-de/insights/decoder/b/behavior-driven-development (Thoughtworks)

---

## Test-Driven Development

TDD simply means: Write tests first! Even before writing the any code.

So we write a test for something that was not yet implemented. And yes, of course that test will fail. This may sound odd at first but TDD follows a simple, iterative cycle known as Red-Green-Refactor:

- **Red**: Write a failing test that describes the desired functionality.
- **Green**: Write the minimal code needed to make the test pass.
- **Refactor**: Improve the code (and tests, if needed) while keeping all tests passing, ensuring the design stays clean.

This cycle ensures that every piece of code is justified by a test, reducing bugs and improving confidence in changes.

## Three Laws of TDD

Robert C. Martin (Uncle Bob) formalized TDD with three key rules:

- You are not allowed to write any production code unless it is to make a failing unit test pass.
- You are not allowed to write any more of a unit test than is sufficient to fail; and compilation failures are failures.
- You are not allowed to write any more production code than is sufficient to pass the currently failing unit test.

## TDD in Action

For a practical example, check out this video of Uncle Bob, where he is coding live, using TDD: https://www.youtube.com/watch?v=rdLO7pSVrMY

It takes time and practice to "master TDD".

## Combine them (TDD + BDD)

TDD and BDD complement each other. TDD ensures your code is correct by driving development through failing tests and the Red-Green-Refactor cycle. BDD ensures your tests focus on what the system should do, not how it does it, by emphasizing behavior over implementation.

Write TDD-style tests to drive small, incremental changes (Red-Green-Refactor). Structure those tests with a BDD mindset, specifying behavior in clear, outcome-focused scenarios.
This approach yields code that is:

- Correct: TDD ensures it works through rigorous testing.
- Maintainable: BDD's focus on behavior keeps tests resilient to implementation changes.
- Well-designed: The discipline of writing tests first encourages modularity, loose coupling, and clear separation of concerns.

## Another Example of BDD

Lastly another example.

Non-BDD:

```java
@Test
public void testHandleMessage() {
    Publisher publisher = new Publisher();
    List<BuilderList> builderLists = publisher.getBuilderLists();
    List<Log> logs = publisher.getLogs();

    Message message = new Message("test");
    publisher.handleMessage(message);

    // Verify build was created
    assertEquals(1, builderLists.size());
    BuilderList lastBuild = getLastBuild(builderLists);
    assertEquals("test", lastBuild.getName());
    assertEquals(2, logs.size());
}
```

With BDD:

```java
@Test
public void shouldGenerateAsyncMessagesFromInterface() {
    Interface messageInterface = Interfaces.createFrom(SimpleMessageService.class);
    PublisherInterface publisher = new PublisherInterface(messageInterface, transport);

    // When we invoke a method on the interface
    SimpleMessageService service = publisher.createPublisher();
    service.sendMessage("Hello");

    // Then a message should be sent through the transport
    verify(transport).send(argThat(message ->
        message.getMethod().equals("sendMessage") &&
        message.getArguments().get(0).equals("Hello")
    ));
}
```

---

# Feedback ⌨️😊

Feel free to contribute by submitting a PR or creating an issue.  
**If this was helpful, you can show support by giving this repository a star! 🌟**

---

# License

MIT
