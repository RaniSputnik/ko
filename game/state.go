package game

import "github.com/RaniSputnik/ko/model"

type Colour byte

const (
	None = Colour(iota)
	Black
	White
)

type State struct {
	stones []Colour
}

func (s State) Stones() []Colour {
	return s.stones
}

// Captives returns the number of stones the given
// player has captured.
func (s State) Captives(colour Colour) int {
	return 0 // TODO
}

// Score gives the total score of the given player.
//
// Score is calculated as 1 point for each empty
// location surrounded by a players stones + 1 pont
// for every stone the player has captured.
func (s State) Score(player *model.User) int {
	return 0 // TODO
}

func newState(m Match) State {
	boardSize := m.Board.Size
	stones := make([]Colour, boardSize*boardSize)
	for _, mv := range m.Board.Moves {
		switch v := mv.(type) {
		case PlayStone:
			i := v.X + v.Y*boardSize

			var pos Colour
			if v.player == m.Owner {
				pos = Black
			} else {
				pos = White
			}

			stones[i] = pos
		}
	}

	return State{stones}
}
