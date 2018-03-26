package game

import (
	"fmt"

	"github.com/RaniSputnik/ko/model"
)

type Colour byte

const (
	None = Colour(iota)
	Black
	White
	// TODO might be good to change this list
	// to add 'Boundary' for when we calculate liberties
)

func (c Colour) String() string {
	switch c {
	case None:
		return "None"
	case Black:
		return "Black"
	case White:
		return "White"
	}
	return fmt.Sprintf("Colour(%d)", c)
}

type State struct {
	boardSize      int
	stones         []Colour
	blackPrisoners int
	whitePrisoners int
}

func (s State) Stones() []Colour {
	return s.stones
}

// Prisoners returns the number of stones of
// the given colour that have been captured.
func (s State) Prisoners(colour Colour) int {
	switch colour {
	case Black:
		return s.blackPrisoners
	case White:
		return s.whitePrisoners
	}
	return 0
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
	blackPrisoners, whitePrisoners := 0, 0
	stones := make([]Colour, boardSize*boardSize)
	for _, mv := range m.Board.Moves {
		switch v := mv.(type) {
		case PlayStone:
			i := v.X + v.Y*boardSize

			// TODO respect 'ColoursReversed' property
			var playerCol Colour
			if v.player == m.Owner {
				playerCol = Black
			} else {
				playerCol = White
			}
			stones[i] = playerCol

			// TODO we will walk the same groups multiple
			// times if we play next to a stone of ours
			// find a way to avoid walking more than once

			walk := [5]pos{
				pos{v.X, v.Y},
				pos{v.X - 1, v.Y},
				pos{v.X, v.Y - 1},
				pos{v.X + 1, v.Y},
				pos{v.X, v.Y + 1},
			}

			groups := [5]group{}
			for i := 0; i < len(walk); i++ {
				groups[i] = findGroup(stones, boardSize, walk[i].x, walk[i].y)
			}

			for _, g := range groups {
				groupSize := len(g.Positions)
				if groupSize == 0 {
					continue
				}

				if g.Liberties == 0 {
					if playerCol == Black {
						whitePrisoners += groupSize
					} else {
						blackPrisoners += groupSize
					}

					for _, i := range g.Positions {
						stones[i] = None // Capture
					}
				}
			}

		}
	}

	return State{boardSize, stones, blackPrisoners, whitePrisoners}
}

func (s State) String() string {
	str := "\n"

	border := "[ ]"
	for i := 0; i < s.boardSize; i++ {
		border += "---"
	}
	border += "[ ]\n"

	i := 0
	str += border
	for y := 0; y < s.boardSize; y++ {
		str += " | "
		for x := 0; x < s.boardSize; x++ {
			var char string
			switch s.stones[i] {
			case Black:
				char = " X "
			case White:
				char = " O "
			default:
				char = " . "
			}
			str += string(char)
			i++
		}
		str += " | \n"
	}
	str += border

	return str
}
