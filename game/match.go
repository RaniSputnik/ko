package game

import (
	"github.com/RaniSputnik/ko/model"
)

type Match struct {
	ID       string
	Owner    *model.User
	Opponent *model.User
	Board    Board

	ColoursReversed bool
}

// Next returns the player who may play next.
func (m Match) Next() *model.User {
	numberOfMovesPlayed := len(m.Board.Moves)
	if m.ColoursReversed {
		numberOfMovesPlayed++
	}

	var nextPlayer *model.User
	if numberOfMovesPlayed%2 == 0 {
		nextPlayer = m.Owner
	} else {
		nextPlayer = m.Opponent
	}

	return nextPlayer
}

var (
	playRules = rules(itMustBeYourTurn, moveMustBeInsideBoardSize)
	skipRules = rules(itMustBeYourTurn)
)

// Play attempts to place a stone from the given player at
// the given position. Returns an error if the move is illegal.
func (m Match) Play(player *model.User, x, y int) (Match, error) {
	mv := PlayStone{player, x, y}

	if err := playRules(m, mv); err != nil {
		return m, err
	}

	m.Board.Moves = append(m.Board.Moves, mv)
	return m, nil
}

// Skip changes the given players turn and passes the turn
// to their opponent.
func (m Match) Skip(player *model.User) (Match, error) {
	mv := Skip{player}

	if err := skipRules(m, mv); err != nil {
		return m, err
	}

	m.Board.Moves = append(m.Board.Moves, mv)
	return m, nil
}

// State returns the current state of the board.
func (m Match) State() State {
	return newState(m)
}
