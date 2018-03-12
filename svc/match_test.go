package svc_test

import (
	"context"
	"testing"

	"github.com/RaniSputnik/ko/kontext"

	"github.com/RaniSputnik/ko/model"
	"github.com/RaniSputnik/ko/svc"
)

func TestJoinMatch(t *testing.T) {
	loggedInUser := model.User{ID: "Bob", Username: "testbob"}
	ctx := kontext.WithUser(context.Background(), loggedInUser)
	mockMatch := model.Match{
		ID:        "test-match",
		Owner:     "Alice",
		BoardSize: 19,
	}
	mockMatchWithOpponent := mockMatch
	mockMatchWithOpponent.Opponent = loggedInUser.ID

	mockStore := &MockStore{}
	mockStore.Func.GetMatch.Returns.Match = mockMatch
	mockStore.Func.SaveMatch.Returns.Match = mockMatchWithOpponent

	m := svc.MatchSvc{Store: mockStore}

	t.Run("CallsGetMatchOnStore", func(t *testing.T) {
		mockStore.Func.GetMatch.WasCalledXTimes = 0
		_, _ = m.JoinMatch(ctx, mockMatch.ID)
		if mockStore.Func.GetMatch.WasCalledXTimes != 1 {
			t.Errorf("Expected 'GetMatch' to be called once but was called '%d' times",
				mockStore.Func.GetMatch.WasCalledXTimes)
		}

		if mockStore.Func.GetMatch.WasCalledWith.MatchID != mockMatch.ID {
			t.Errorf("Expected 'GetMatch' to be called with matchID: '%s', but instead got '%s'",
				mockMatch.ID, mockStore.Func.GetMatch.WasCalledWith.MatchID)
		}
	})

	t.Run("ReturnsValidMatch", func(t *testing.T) {
		gotMatch, gotErr := m.JoinMatch(ctx, mockMatch.ID)
		if gotErr != nil {
			t.Errorf("Expected nil error, got '%s'", gotErr)
		}
		if gotMatch.ID != mockMatch.ID {
			t.Errorf("Expected id: '%s', got: '%s'", mockMatch.ID, gotMatch.ID)
		}
		if gotMatch.Owner != mockMatch.Owner {
			t.Errorf("Expected owner: '%s', got: '%s'", mockMatch.Owner, gotMatch.Owner)
		}
	})

	t.Run("SetsOpponentField", func(t *testing.T) {
		gotMatch, _ := m.JoinMatch(ctx, mockMatch.ID)
		if gotMatch.Opponent != loggedInUser.ID {
			t.Errorf("Expected opponent id: '%s', got: '%s'", loggedInUser.ID, gotMatch.Opponent)
		}
	})

	t.Run("CallsSaveMatchOnStore", func(t *testing.T) {
		mockStore.Func.SaveMatch.WasCalledXTimes = 0
		_, _ = m.JoinMatch(ctx, mockMatch.ID)
		if mockStore.Func.SaveMatch.WasCalledXTimes != 1 {
			t.Errorf("Expected 'SaveMatch' to be called once but was called '%d' times",
				mockStore.Func.SaveMatch.WasCalledXTimes)
		}

		if mockStore.Func.SaveMatch.WasCalledWith.Match != mockMatchWithOpponent {
			t.Errorf("Expected 'SaveMatch' to be called with match: '%v', but instead got '%v'",
				mockMatchWithOpponent, mockStore.Func.SaveMatch.WasCalledWith.Match)
		}
	})

	t.Run("FailsWhenUserJoinsTheirOwnMatch", func(t *testing.T) {
		ownMatch := model.Match{
			ID:        "our-own-match",
			Owner:     loggedInUser.ID,
			BoardSize: 19,
		}
		mockStore := &MockStore{}
		mockStore.Func.GetMatch.Returns.Match = ownMatch
		m := svc.MatchSvc{Store: mockStore}

		_, err := m.JoinMatch(ctx, ownMatch.ID)
		if _, ok := err.(model.ErrJoinedOwnMatch); !ok {
			t.Errorf("Expected error of type: 'ErrJoinedOwnMatch', but got: '%v'", err)
		}
	})

	t.Run("FailsWhenTheMatchIsNotFound", func(t *testing.T) {
		mockStore := &MockStore{}
		mockStore.Func.GetMatch.Returns.Err = model.ErrMatchNotFound{}
		m := svc.MatchSvc{Store: mockStore}

		_, err := m.JoinMatch(ctx, mockMatch.ID)
		if _, ok := err.(model.ErrMatchNotFound); !ok {
			t.Errorf("Expected error of type: 'ErrMatchNotFound', but got: '%v'", err)
		}
	})

	t.Run("FailsWhenMatchAlreadyHasAnOpponent", func(t *testing.T) {
		mockStore := &MockStore{}
		mockStore.Func.GetMatch.Returns.Match = mockMatchWithOpponent
		m := svc.MatchSvc{Store: mockStore}

		_, err := m.JoinMatch(ctx, mockMatch.ID)
		if _, ok := err.(model.ErrMatchAlreadyFull); !ok {
			t.Errorf("Expected error of type: 'ErrMatchAlreadyFull', but got: '%v'", err)
		}
	})
}
