package game

import "github.com/RaniSputnik/ko/model"

type playRule func(m Match, x, y int) error

func moveMustBeInsideBoardSize(m Match, x, y int) error {
	boardSize := m.Board.Size
	if x < 0 || y < 0 || x >= boardSize || y >= boardSize {
		return model.ErrOutOfBounds{
			X:         x,
			Y:         y,
			BoardSize: boardSize,
		}
	}
	return nil
}
