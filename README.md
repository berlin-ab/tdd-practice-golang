# TDD Practice

Practice doing test driven developemnt by following the Red -> Green -> Refactor testing loop.

Make a red test, make the simplest solution to make the test green, make code improvements while keeping the tests green

### Rock Paper Scissors

Test all the possible outcomes of a single round of Rock Paper Scissors. For example:

* `play(Rock, Scissors)` should have the result `PlayerOneWins`
* `play(Rock, Rock)` should have the result `TieGame`

*Hint: a Gomega Matcher called `HaveTheResult` exists to translate enum values into human readable results*

To run test suite:

```
# single run
make test-rock-paper-scissors
```


### Supermarket Kata

loosely based on:
 
http://codekata.com/kata/kata01-supermarket-pricing/

The goal of this exercise is to introduce the concept of dependency inversion. Isolate the things that change from the things that stay constant.

Goals:

* Allow the inventory to add new inventory items
* Add discount to specific items
* Add new tests while attempting to keep the existing tests passing (note: requires some modification to the existing tests)

# single run
```
make test-supermarket
```

### Lightswitch Controller

Your team is responsible for developing a lightswitch controller for a smarthome. Your boss wants you to integrate with their preferred vendor (Edison), but doesn't want to pay for the license to use the Edison SDK. Your boss still wants your code to work ASAP.

After creating an initial implementation, your team has realized that it is hard to test. 

Through writing software, answer these questions:

* How can we test the logic of the lightswitch controller without paying for the edison lightbulb license?
* How should we organize the packages of our software?

Introducing the *Dependency Inversion Principle:*

https://en.wikipedia.org/wiki/Dependency_inversion_principle

- High-level modules should not depend on low-level modules. Both should depend on abstractions (e.g., interfaces).
- Abstractions should not depend on details. Details (concrete implementations) should depend on abstractions.