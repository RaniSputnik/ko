package game

import "github.com/RaniSputnik/ko/model"

type playRule func(m Match, mv PlayStone) error

func itMustBeYourTurn(m Match, mv PlayStone) error {
	if next := m.Next(); next != mv.Player() {
		return model.ErrNotYourTurn{Next: next}
	}
	return nil
}

func moveMustBeInsideBoardSize(m Match, mv PlayStone) error {
	boardSize := m.Board.Size
	if mv.X < 0 || mv.Y < 0 || mv.X >= boardSize || mv.Y >= boardSize {
		return model.ErrOutOfBounds{
			X:         mv.X,
			Y:         mv.Y,
			BoardSize: boardSize,
		}
	}
	return nil
}
