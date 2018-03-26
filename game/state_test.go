package game_test

import (
	"testing"

	"github.com/RaniSputnik/ko/game"
	"github.com/RaniSputnik/ko/model"
)

func TestStateColourString(t *testing.T) {
	testCases := []struct {
		Test   game.Colour
		Expect string
	}{
		{Test: game.Black, Expect: "Black"},
		{Test: game.White, Expect: "White"},
		{Test: game.None, Expect: "None"},
		{Test: game.Colour(4), Expect: "Colour(4)"},
		{Test: game.Colour(7), Expect: "Colour(7)"},
	}

	for _, test := range testCases {
		if got := test.Test.String(); got != test.Expect {
			t.Errorf("Expected: %s, Got: %s", test.Expect, got)
		}
	}
}

func TestStateReturnsStones(t *testing.T) {
	boardSize := 9

	testCases := []struct {
		Desc   string
		Match  game.Match
		Expect []game.Colour
	}{
		{
			Desc:   "One move should result in one stone",
			Match:  matchWithMoves(boardSize, pos{0, 0}),
			Expect: state(boardSize, stone{game.Black, 0, 0}),
		},
		{
			Desc:  "Three stones should be alternating in colour",
			Match: matchWithMoves(boardSize, pos{0, 0}, pos{1, 2}, pos{5, 4}),
			Expect: state(boardSize,
				stone{game.Black, 0, 0},
				stone{game.White, 1, 2},
				stone{game.Black, 5, 4},
			),
		},
	}

	for _, test := range testCases {
		got := test.Match.State().Stones()
		if len(got) != len(test.Expect) {
			t.Errorf("%s. Expected '%d' positions, Got: '%d'", test.Desc, len(test.Expect), len(got))
		}
		for i, gotStone := range got {
			if gotStone != test.Expect[i] {
				t.Errorf("%s. Expected stone '%d' to be: %v, Got: %v", test.Desc, i+1, test.Expect, got)
			}
		}
	}
}

func TestStonesWithoutLibertiesAreCaptured(t *testing.T) {
	boardSize := 9

	testCases := []struct {
		Name           string
		Match          game.Match
		Captures       []pos
		BlackPrisoners int
		WhitePrisoners int
	}{
		{
			Name: "FullSurround",
			Match: matchWithMoves(game.BoardSizeTiny,
				pos{2, 2}, pos{3, 2}, pos{3, 3}, pos{8, 8}, pos{4, 2}, pos{8, 7}, pos{3, 1}),
			Captures:       []pos{pos{3, 2}},
			BlackPrisoners: 0,
			WhitePrisoners: 1,
		},
		{
			Name: "TopLeftCorner",
			Match: matchWithMoves(game.BoardSizeTiny,
				pos{0, 0}, pos{1, 0}, pos{8, 8}, pos{0, 1}),
			Captures:       []pos{pos{0, 0}},
			BlackPrisoners: 1,
			WhitePrisoners: 0,
		},
		{
			Name: "BottomRightCorner",
			Match: matchWithMoves(game.BoardSizeTiny,
				pos{8, 8}, pos{7, 8}, pos{0, 0}, pos{8, 7}),
			Captures:       []pos{pos{8, 8}},
			BlackPrisoners: 1,
			WhitePrisoners: 0,
		},
		{
			Name: "RightHandSide",
			Match: matchWithMoves(game.BoardSizeTiny,
				pos{8, 3}, pos{8, 4}, pos{7, 4}, pos{8, 8}, pos{8, 5}),
			Captures:       []pos{pos{8, 4}},
			BlackPrisoners: 0,
			WhitePrisoners: 1,
		},
		{
			Name: "SurroundedGroup",
			Match: matchWithMoves(game.BoardSizeTiny,
				pos{3, 6}, pos{3, 5}, pos{2, 5}, pos{3, 4}, pos{4, 5},
				pos{4, 4}, pos{2, 4}, pos{4, 3}, pos{3, 3}, pos{8, 8},
				pos{5, 3}, pos{8, 7}, pos{5, 4}, pos{8, 6}, pos{4, 2},
			),
			Captures:       []pos{pos{3, 4}, pos{3, 5}, pos{4, 3}, pos{4, 4}},
			BlackPrisoners: 0,
			WhitePrisoners: 4,
		},
		{
			Name: "GroupAgainstWall",
			Match: matchWithMoves(game.BoardSizeTiny,
				pos{0, 6}, pos{0, 5}, pos{1, 5}, pos{0, 4}, pos{1, 3},
				pos{1, 4}, pos{0, 2}, pos{0, 3}, pos{2, 4},
			),
			Captures:       []pos{pos{0, 5}, pos{0, 4}, pos{1, 4}, pos{0, 3}},
			BlackPrisoners: 0,
			WhitePrisoners: 4,
		},
		// TODO test donut group
	}

	for i, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			state := test.Match.State()
			stones := state.Stones()
			t.Logf("Test case #%d: %s", i+1, state)

			// TODO use a helper method instead
			for _, capture := range test.Captures {
				index := capture.X + capture.Y*boardSize
				if got := stones[index]; got != game.None {
					t.Errorf("Expected stone at %d,%d to have been captured.",
						capture.X, capture.Y)
				}
			}

			if got := state.Prisoners(game.Black); got != test.BlackPrisoners {
				t.Errorf("Expected: '%d' black stone(s) to be captured, got: '%d'", test.BlackPrisoners, got)
			}
			if got := state.Prisoners(game.White); got != test.WhitePrisoners {
				t.Errorf("Expected: '%d' white stone(s) to be captured, got: '%d'", test.WhitePrisoners, got)
			}
		})
	}
}

func state(boardSize int, stones ...stone) []game.Colour {
	s := make([]game.Colour, boardSize*boardSize)
	for _, st := range stones {
		i := st.X + st.Y*boardSize
		s[i] = st.Colour
	}
	return s
}

func matchWithMoves(boardSize int, moves ...pos) game.Match {
	m := game.Match{
		Owner:    &Alice,
		Opponent: &Bob,
		Board: game.Board{
			Size: boardSize,
		},
	}

	for i, mv := range moves {
		var player *model.User
		if i%2 == 0 {
			player = &Alice
		} else {
			player = &Bob
		}

		var err error
		if m, err = m.Play(player, mv.X, mv.Y); err != nil {
			panic(err)
		}
	}

	return m
}
