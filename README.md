# TDD Practice

### Rock Paper Scissors

Test all the possible outcomes of a single round of Rock Paper Scissors. For example:

* `play(Rock, Scissors)` should have the result `PlayerOneWins`
* `play(Rock, Rock)` should have the result `TieGame`

*Hint: a Gomega Matcher called `HaveTheResult` exists to translate enum values into human readable results*

To run test suite:

```
# single run
make test-rock-paper-scissors

# run continuously
make watch-rock-paper-scissors
```