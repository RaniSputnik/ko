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
			var pos Colour
			if v.player == m.Owner {
				pos = Black
			} else {
				pos = White
			}
			stones[i] = pos

			var captured int
			stones, captured = recalculateLiberties(stones, boardSize, v.X, v.Y)
			if pos == Black {
				whitePrisoners += captured
			} else {
				blackPrisoners += captured
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
