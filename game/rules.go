package game

import (
	"github.com/RaniSputnik/ko/model"
)

type ruleFunc func(m Match, mv Move) error

func rules(rules ...ruleFunc) ruleFunc {
	// No rules? Do nothing
	if len(rules) == 0 {
		return func(m Match, mv Move) error { return nil }
	}

	// Only one rule, then return just that rule
	if len(rules) == 1 {
		return rules[0]
	}

	// Otherwise a func that runs all the rules
	// returning as soon as an error is encountered
	return func(m Match, mv Move) error {
		for _, rule := range rules {
			if err := rule(m, mv); err != nil {
				return err
			}
		}
		return nil
	}
}

func theGameMustBeInProgress(m Match, mv Move) error {
	if m.Opponent == nil {
		return model.ErrMatchNotStarted{}
	}
	// TODO nobody has resigned and the game is not over
	return nil
}

func playerMustBeInGame(m Match, mv Move) error {
	player := mv.Player()
	if m.Owner != player && m.Opponent != player {
		return model.ErrNotParticipating{}
	}
	return nil
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

func positionMustNotBeOccupied(m Match, mv Move) error {
	// TODO pass this to rule func? If many rules need this
	// will be inefficient to recalculate it each time
	state := m.State()
	ps := mv.(PlayStone)

	stones := state.Stones()
	i := ps.X + ps.Y*m.Board.Size
	if stones[i] != None {
		return model.ErrPositionOccupied{}
	}
	return nil
}
