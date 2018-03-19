package game

import (
	"github.com/RaniSputnik/ko/model"
)

type Match struct {
	ID       string
	Owner    *model.User
	Opponent *model.User
	Board    Board
}

type Board struct {
	Size  int
	Moves []Move
}

type Stone struct {
	Player *model.User
	X, Y   int
}

func (m Match) Play(user *model.User, x, y int) (Match, error) {
	playRules := []playRule{
		moveMustBeInsideBoardSize,
	}

	for _, rule := range playRules {
		if err := rule(m, x, y); err != nil {
			return m, err
		}
	}

	mv := PlayStone{user, x, y}

	m.Board.Moves = append(m.Board.Moves, mv)
	return m, nil
}
