package game_test

import (
	"testing"

	"github.com/RaniSputnik/ko/game"
	"github.com/RaniSputnik/ko/model"
)

const TestID = "test123"

var Alice = model.User{ID: "Alice", Username: "testalice"}
var Bob = model.User{ID: "Bob", Username: "testbob"}

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
			ID:       TestID,
			Owner:    &Alice,
			Opponent: &Bob,
			Board: game.Board{
				Size:  boardSize,
				Moves: []game.Move{},
			},
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
