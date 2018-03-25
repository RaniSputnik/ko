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

func TestAStoneWithoutLibertiesIsCaptured(t *testing.T) {
	boardSize := 9

	testCases := []struct {
		Name               string
		Match              game.Match
		CaptureX, CaptureY int
		BlackPrisoners     int
		WhitePrisoners     int
	}{
		{
			Name: "FullSurround",
			Match: matchWithMoves(game.BoardSizeTiny,
				pos{2, 2}, pos{3, 2}, pos{3, 3}, pos{8, 8}, pos{4, 2}, pos{8, 7}, pos{3, 1}),
			CaptureX:       3,
			CaptureY:       2,
			BlackPrisoners: 0,
			WhitePrisoners: 1,
		},
		{
			Name: "TopLeftCorner",
			Match: matchWithMoves(game.BoardSizeTiny,
				pos{0, 0}, pos{1, 0}, pos{8, 8}, pos{0, 1}),
			CaptureX:       0,
			CaptureY:       0,
			BlackPrisoners: 1,
			WhitePrisoners: 0,
		},
		{
			Name: "BottomRightCorner",
			Match: matchWithMoves(game.BoardSizeTiny,
				pos{8, 8}, pos{7, 8}, pos{0, 0}, pos{8, 7}),
			CaptureX:       8,
			CaptureY:       8,
			BlackPrisoners: 1,
			WhitePrisoners: 0,
		},
		{
			Name: "RightHandSide",
			Match: matchWithMoves(game.BoardSizeTiny,
				pos{8, 3}, pos{8, 4}, pos{7, 4}, pos{8, 8}, pos{8, 5}),
			CaptureX:       8,
			CaptureY:       4,
			BlackPrisoners: 0,
			WhitePrisoners: 1,
		},
	}

	for i, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			state := test.Match.State()
			stones := state.Stones()
			t.Logf("Test case #%d: %s", i+1, state)

			// TODO use a helper method instead
			index := test.CaptureX + test.CaptureY*boardSize
			if got := stones[index]; got != game.None {
				t.Errorf("Expected stone at %d,%d to have been captured.",
					test.CaptureX, test.CaptureY)
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
			//t.Fatalf("Failed to setup test: %s", err)
			panic(err)
		}
	}

	return m
}
