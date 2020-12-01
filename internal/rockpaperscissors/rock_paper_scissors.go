package rockpaperscissors

type Result int

const (
	PlayerOneWins Result = iota
	PlayerTwoWins
	TieGame
	InvalidGame
)

type option int

const (
	Rock option = iota
	Paper
	Scissors
)

func Play(playerOne, playerTwo option) Result {
	return InvalidGame
}