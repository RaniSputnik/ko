package svc_test

import (
	"context"
	"testing"

	"github.com/RaniSputnik/ko/kontext"
	"github.com/RaniSputnik/ko/model"
	"github.com/RaniSputnik/ko/svc"
)

func TestPlay(t *testing.T) {
	loggedInUser := Bob
	ctx := kontext.WithUser(context.Background(), loggedInUser)

	mockStore := &MockStore{}
	p := svc.PlaySvc{MoveStore: mockStore}

	mockMatch := model.Match{
		ID:        "test-match",
		Owner:     "Alice",
		BoardSize: 19,
	}

	t.Run("ReturnsValidPlayStoneEvent", func(t *testing.T) {
		playX, playY := 1, 2
		ev, err := p.Play(ctx, mockMatch.ID, playX, playY)

		if err != nil {
			t.Errorf("Expected nil error, got: '%v'", err)
		}
		if ev.PlayerID != loggedInUser.ID {
			t.Errorf("Expected PlayerID: '%s', Got: '%s'", loggedInUser.ID, ev.PlayerID)
		}
		if ev.X != playX {
			t.Errorf("Expected X: '%d', Got: '%d'", playX, ev.X)
		}
		if ev.X != playX {
			t.Errorf("Expected Y: '%d', Got: '%d'", playY, ev.Y)
		}
	})

	t.Run("CallsSaveMoveOnStore", func(t *testing.T) {
		playX, playY := 1, 2
		mockStore.Func.SaveMove.WasCalledXTimes = 0

		p.Play(ctx, mockMatch.ID, playX, playY)

		if mockStore.Func.SaveMove.WasCalledXTimes != 1 {
			t.Errorf("Expected save move to be called once but instead was called '%d' times.",
				mockStore.Func.SaveMove.WasCalledXTimes)
		}
	})
}
