package svc_test

import (
	"context"

	"github.com/RaniSputnik/ko/model"
)

var Alice = model.User{}
var Bob = model.User{ID: "Bob", Username: "testbob"}

type MockStore struct {
	Func struct {
		SaveMatch struct {
			WasCalledXTimes int
			WasCalledWith   struct {
				Ctx   context.Context
				Match model.Match
			}
			Returns struct {
				Match model.Match
				Err   error
			}
		}
		GetMatches struct {
			WasCalledXTimes int
			WasCalledWith   struct {
				Ctx    context.Context
				UserID string
			}
			Returns struct {
				Matches []model.Match
				Err     error
			}
		}
		GetMatch struct {
			WasCalledXTimes int
			WasCalledWith   struct {
				Ctx     context.Context
				MatchID string
			}
			Returns struct {
				Match model.Match
				Err   error
			}
		}
	}
}

func (m *MockStore) SaveMatch(ctx context.Context, match model.Match) (model.Match, error) {
	m.Func.SaveMatch.WasCalledWith.Ctx = ctx
	m.Func.SaveMatch.WasCalledWith.Match = match
	m.Func.SaveMatch.WasCalledXTimes++
	return m.Func.SaveMatch.Returns.Match, m.Func.SaveMatch.Returns.Err
}

func (m *MockStore) GetMatches(ctx context.Context, userID string) ([]model.Match, error) {
	m.Func.GetMatches.WasCalledWith.Ctx = ctx
	m.Func.GetMatches.WasCalledWith.UserID = userID
	m.Func.GetMatches.WasCalledXTimes++
	return m.Func.GetMatches.Returns.Matches, m.Func.GetMatches.Returns.Err
}

func (m *MockStore) GetMatch(ctx context.Context, matchID string) (model.Match, error) {
	m.Func.GetMatch.WasCalledWith.Ctx = ctx
	m.Func.GetMatch.WasCalledWith.MatchID = matchID
	m.Func.GetMatch.WasCalledXTimes++
	return m.Func.GetMatch.Returns.Match, m.Func.GetMatch.Returns.Err
}
