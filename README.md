# Trip Service Kata

Kata for legacy code hands-on session. The objective is to test and refactor the legacy TripService class.

The end result should be well-crafted code that express the domain.

`TripService.getTripsByUser` is a use case to return a list of trips from your single friend.

## Lesson 1 - Extract and Override

`TripService.getTripsByUser` is a use case to return a list of trips from your single friend.

Extract all the dependencies with side effects related code to make `TripService.getTripsByUser` testable.

✅ Expected output:

- Pass all tests:
  - It should get a list of trips from a user if the user is your friend
  - It should not get a list of trips from a user if the user is not your friend
  - It should throw UserNotLoggedInException if you are not logged in
- Two extra protected methods to in `TripService`

<details><summary>Tips in case you need</summary>
<p>
1. Highlight the part to `getLoggedUser` and do `Extract Method` to replace with a new protected method
2. Highlight the part to `findTripsByUser` and do `Extract Method` to replace with a new protected method
3. Create a test file
4. In the test file, create a child class of `TripService` and override the two protected methods mentioned above to return fixed data
5. In the test file, create some test cases with some fake data.
</p>
</details>

## Lesson 2 - Extract Interface

Extract Interface from the dependencies with side effects of `TripService` to make `TripService.getTripsByUser` testable and detectable (able to be sensed)

✅ Expected output:

- Pass all tests:
  - It should get a list of trips from a user if the user is your friend
  - It should not get a list of trips from a user if the user is not your friend
  - It should throw UserNotLoggedInException if you are not logged in
- Two extra protected methods to in `TripService`

<details><summary>Tips in case you need</summary>
<p>
1. Create an interface `IUserSession` and put methods you need in `UserSession` into it.
2. Make the `UserSession` implements the interface
3. Create an interface `ITripDAO` and put methods you need in `TripDAO` into it.
4. Make the `TripDAO` implements the interface
5. Create a test file
6. In the test file, create subclasses of the two interfaces mentioned above and return fixed data.
7. In the test file, create a test case to make it fail
8. In the `TripService`, dependency inject the two modules into the constructor of the `TripService`
9. In the `TripService`, change the dependency modules to the interface.
10. Go back to the tests, pass the subclasses you created in step 6 to the test cases and make tests pass.
</p>
</details>

## Lesson 3 - Make changes

For privacy concerns, we want to add an attribute `public` to `Trip` which default is true. If the `trip.public` is false, then even your friends cannot access this trip by `TripService.getTripsByUser`.

Example:

```gherkin
  Scenario: You cannot access trips that was not public event though the owner is your friend
    Given Foo is Bar's friend
    And Bar has trips:
      | trip | public |
      | Ta   | true   |
      | Tb   | false  |
    When Bar request Bar's trip list
    Then it should display Ta
```

## Retrospective

- Discuss the difference between extract and override and extract interface
  - Why to use
  - When to use
- Did you use the IDE refactor tools? How do you think about them?
- Does tests help you build confidence to make changes to code?
- Is there any possible to improve your tests in readability and maintainability?

## Advanced Lessons

## Notes

1. Use refactor tools as many as you can to prevent manually modification.
2. Make sure your results and changes are covered by tests
