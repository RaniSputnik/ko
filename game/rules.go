package game

import "github.com/RaniSputnik/ko/model"

type ruleFunc func(m Match, mv Move) error

func rules(rules ...ruleFunc) ruleFunc {
	return func(m Match, mv Move) error {
		for _, rule := range rules {
			if err := rule(m, mv); err != nil {
				return err
			}
		}
		return nil
	}
}

func itMustBeYourTurn(m Match, mv Move) error {
	if next := m.Next(); next != mv.Player() {
		return model.ErrNotYourTurn{Next: next}
	}
	return nil
}

func moveMustBeInsideBoardSize(m Match, mv Move) error {
	ps := mv.(PlayStone)

	boardSize := m.Board.Size
	if ps.X < 0 || ps.Y < 0 || ps.X >= boardSize || ps.Y >= boardSize {
		return model.ErrOutOfBounds{
			X:         ps.X,
			Y:         ps.Y,
			BoardSize: boardSize,
		}
	}
	return nil
}
