package rockpaperscissors_test

import (
	"fmt"

	"github.com/onsi/gomega/types"

	"tddpractice/internal/rockpaperscissors"
)

type matchResultMatcher struct {
	expected rockpaperscissors.Result
}

var resultToString = map[rockpaperscissors.Result]string{
	rockpaperscissors.PlayerOneWins: "Player One Wins",
	rockpaperscissors.PlayerTwoWins: "Player Two Wins",
	rockpaperscissors.TieGame: "Tie Game",
	rockpaperscissors.InvalidGame: "Invalid Game",
}

func (m matchResultMatcher) Match(actual interface{}) (success bool, err error) {
	actual, ok := actual.(rockpaperscissors.Result)
	if !ok {
		return false, fmt.Errorf("failed to convert into result")
	}

	return m.expected == actual, nil
}

func (m matchResultMatcher) FailureMessage(actual interface{}) (message string) {
	actualResult, _ := actual.(rockpaperscissors.Result)

	return fmt.Sprintf(`Expected "%v" to be the result of the game, got "%v"`,
		resultToString[m.expected],
		resultToString[actualResult])
}

func (m matchResultMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	actualResult, _ := actual.(rockpaperscissors.Result)

	return fmt.Sprintf(`Expected "%v" to NOT be the result of the game, Got "%v".`,
		resultToString[m.expected],
		resultToString[actualResult])
}

func HaveTheResult(result rockpaperscissors.Result) (types.GomegaMatcher) {
	return matchResultMatcher{
		expected: result,
	}
}

