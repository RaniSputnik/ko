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

type Move struct {
	Player *model.User
	X, Y   int
}

func (m Match) Play(user *model.User, x, y int) (Match, error) {
	mv := Move{
		Player: user,
		X:      x,
		Y:      y,
	}

	m.Board.Moves = append(m.Board.Moves, mv)
	return m, nil
}
