package game_test

import (
	"testing"

	"github.com/RaniSputnik/ko/game"
	"github.com/RaniSputnik/ko/model"
)

const TestID = "test123"

var Alice = model.User{ID: "Alice", Username: "testalice"}
var Bob = model.User{ID: "Bob", Username: "testbob"}
var Clive = model.User{ID: "Clive", Username: "testclive"}

func TestPlayAddsAMove(t *testing.T) {
	loggedInUser := &Alice

	m := game.Match{
		ID:       TestID,
		Owner:    &Alice,
		Opponent: &Bob,
		Board: game.Board{
			Size:  game.BoardSizeNormal,
			Moves: []game.Move{},
		},
	}

	playX, playY := 3, 2
	match, err := m.Play(loggedInUser, playX, playY)

	if err != nil {
		t.Errorf("Expected '<nil>' error, got '%v'", err)
	}

	if len(match.Board.Moves) != 1 {
		t.Errorf("Expected '1' move, got '%v' move(s)", match.Board.Moves)
	}
}

func TestCanNotPlayOutOfBounds(t *testing.T) {
	loggedInUser := &Alice

	newMatch := func(boardSize int) game.Match {
		return game.Match{
			Owner:    &Alice,
			Opponent: &Bob,
			Board:    game.Board{Size: boardSize},
		}
	}

	testCases := []struct{ PlayX, PlayY, BoardSize int }{
		{PlayX: -1, PlayY: 0, BoardSize: 9},
		{PlayX: -3, PlayY: -3, BoardSize: 9},
		{PlayX: 0, PlayY: 19, BoardSize: 19},
		{PlayX: 5, PlayY: -3, BoardSize: 9},
		{PlayX: 13, PlayY: 13, BoardSize: 9},
		{PlayX: 5, PlayY: 12, BoardSize: 9},
		{PlayX: 15, PlayY: 3, BoardSize: 4},
		{PlayX: 19, PlayY: 3, BoardSize: 19},
		{PlayX: 32, PlayY: 400, BoardSize: 19},
	}

	for _, test := range testCases {
		m := newMatch(test.BoardSize)
		match, err := m.Play(loggedInUser, test.PlayX, test.PlayY)

		expected := model.ErrOutOfBounds{X: test.PlayX, Y: test.PlayY, BoardSize: test.BoardSize}
		if err != expected {
			t.Errorf("Expected '%+v' error, got '%+v'", expected, err)
		}

		if len(match.Board.Moves) > 0 {
			t.Errorf("Expected '0' moves, got '%v'", match.Board.Moves)
		}
	}
}

func TestNextReturnsNextUsersTurn(t *testing.T) {
	genMoves := func(n int) []game.Move {
		moves := make([]game.Move, n)
		for i := 0; i < n; i++ {
			moves[i] = mockMove{}
		}
		return moves
	}

	testCases := []struct {
		Match  game.Match
		Expect *model.User
	}{
		{
			Match: game.Match{
				Owner:    &Alice,
				Opponent: &Bob,
			},
			Expect: &Alice,
		},
		{
			Match: game.Match{
				Owner:    &Bob,
				Opponent: &Alice,
			},
			Expect: &Bob,
		},
		{
			Match: game.Match{
				Owner:           &Bob,
				Opponent:        &Alice,
				ColoursReversed: true,
			},
			Expect: &Alice,
		},
		{
			Match: game.Match{
				Owner:    &Alice,
				Opponent: &Bob,
				Board:    game.Board{Moves: genMoves(1)},
			},
			Expect: &Bob,
		},
		{
			Match: game.Match{
				Owner:           &Alice,
				Opponent:        &Bob,
				Board:           game.Board{Moves: genMoves(2)},
				ColoursReversed: true,
			},
			Expect: &Bob,
		},
		{
			Match: game.Match{
				Owner:    &Alice,
				Opponent: &Bob,
				Board:    game.Board{Moves: genMoves(15)},
			},
			Expect: &Bob,
		},
	}

	for i, test := range testCases {
		if got := test.Match.Next(); got != test.Expect {
			t.Errorf("Match no.%d, Expected next player: %v, Got: %v", i+1, test.Expect, got)
		}
	}
}

func TestOnlyTheNextPlayerCanPlay(t *testing.T) {
	m := game.Match{
		Owner:    &Alice,
		Opponent: &Bob,
		Board:    game.Board{Size: game.BoardSizeNormal},
	}

	// Alice is the owner, it is her turn
	_, err := m.Play(&Bob, 0, 0)
	expected := model.ErrNotYourTurn{Next: &Alice}
	if err != expected {
		t.Errorf("Match where owner starts. Expected: %v, got: %v", expected, err)
	}

	// Reverse the colours, it should be Bob's turn
	m.ColoursReversed = true

	m, err = m.Play(&Alice, 0, 0)
	expected = model.ErrNotYourTurn{Next: &Bob}
	if err != expected {
		t.Errorf("Match with colours reversed. Expected: %v, got: %v", expected, err)
	}

	// Add one move now it's Alice's turn
	m, _ = m.Play(&Bob, 0, 0)

	m, err = m.Play(&Bob, 1, 1)
	expected = model.ErrNotYourTurn{Next: &Alice}
	if err != expected {
		t.Errorf("Match with 1 move. Expected: %v, got: %v", expected, err)
	}
}

func TestCanNotMoveWhenTheGameHasNotStarted(t *testing.T) {
	m := game.Match{
		Owner: &Alice,
		Board: game.Board{Size: game.BoardSizeNormal},
	}

	t.Run("Play", func(t *testing.T) {
		_, err := m.Play(&Alice, 0, 0)
		expectError(t, model.ErrMatchNotStarted{}, err)
	})

	t.Run("Skip", func(t *testing.T) {
		_, err := m.Skip(&Alice)
		expectError(t, model.ErrMatchNotStarted{}, err)
	})

	/*t.Run("Resign", func(t *testing.T) {
		_, err := m.Resign(&Alice)
		expectError(t, model.ErrMatchNotStarted{}, err)
	})*/
}

func TestCanNotMoveWhenNotPartOfGame(t *testing.T) {
	m := game.Match{
		Owner:    &Alice,
		Opponent: &Bob,
		Board:    game.Board{Size: game.BoardSizeNormal},
	}

	t.Run("Play", func(t *testing.T) {
		_, err := m.Play(&Clive, 0, 0)
		expectError(t, model.ErrNotParticipating{}, err)
	})

	t.Run("Skip", func(t *testing.T) {
		_, err := m.Skip(&Clive)
		expectError(t, model.ErrNotParticipating{}, err)
	})
}

func TestCannotSkipWhenItIsNotYourTurn(t *testing.T) {
	m := game.Match{
		Owner:    &Alice,
		Opponent: &Bob,
	}

	_, err := m.Skip(&Bob)
	expectError(t, model.ErrNotYourTurn{Next: &Alice}, err)
}

func TestSkipChangesTurn(t *testing.T) {
	m, err := game.Match{
		Owner:    &Alice,
		Opponent: &Bob,
	}.Skip(&Alice)

	if err != nil {
		t.Errorf("Expected: '<nil>' error when skipping, got: '%v'", err)
	}

	if next := m.Next(); next != &Bob {
		t.Errorf("Expected next: '%v', got: '%v'", &Bob, next)
	}

	m.ColoursReversed = true

	if next := m.Next(); next != &Alice {
		t.Errorf("Expected next: '%v', got: '%v'", &Alice, next)
	}
}

type pos struct {
	X, Y int
}

type stone struct {
	Colour game.Colour
	X, Y   int
}

type mockMove struct{}

func (mv mockMove) String() string { return "A mock move" }

func (mv mockMove) Player() *model.User { return nil }

func expectError(t *testing.T, expected, got error) {
	if got != expected {
		t.Errorf("Expected: '%v', got: '%v", expected, got)
	}
}
