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
