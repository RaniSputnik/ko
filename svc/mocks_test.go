package svc_test

import (
	"context"

	"github.com/RaniSputnik/ko/data"

	"github.com/RaniSputnik/ko/model"
)

var (
	Alice = model.User{ID: "Alice", Username: "testalice"}
	Bob   = model.User{ID: "Bob", Username: "testbob"}
)

const (
	MatchID12345 = "TWF0Y2g6MTIzNDU="
	MatchID67890 = "TWF0Y2g6Njc4OTA="
)

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
		SaveMove struct {
			WasCalledXTimes int
			WasCalledWith   struct {
				Ctx  context.Context
				Move data.Move
			}
			Returns struct {
				Move data.Move
				Err  error
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

func (m *MockStore) SaveMove(ctx context.Context, move data.Move) (data.Move, error) {
	m.Func.SaveMove.WasCalledWith.Ctx = ctx
	m.Func.SaveMove.WasCalledWith.Move = move
	m.Func.SaveMove.WasCalledXTimes++
	return m.Func.SaveMove.Returns.Move, m.Func.GetMatch.Returns.Err
}
