package svc

import (
	"context"

	"github.com/RaniSputnik/ko/data"
	"github.com/RaniSputnik/ko/kontext"
	"github.com/RaniSputnik/ko/model"
)

const (
	BoardSizeTiny   = 9
	BoardSizeSmall  = 13
	BoardSizeNormal = 19
)

type MatchSvc struct {
	Store data.MatchStore
}

func (svc MatchSvc) CreateMatch(ctx context.Context, boardSize int) (model.Match, error) {
	user := kontext.MustGetUser(ctx)
	match := model.Match{Owner: user.ID, BoardSize: boardSize}
	return svc.Store.SaveMatch(ctx, match)
}

func (svc MatchSvc) GetMatches(ctx context.Context) ([]model.Match, error) {
	user := kontext.MustGetUser(ctx)
	matches, err := svc.Store.GetMatches(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	for i, match := range matches {
		match.ID = model.EncodeID(model.KindMatch, match.ID)
		matches[i] = match
	}
	return matches, nil
}

func (svc MatchSvc) JoinMatch(ctx context.Context, matchID string) (model.Match, error) {
	user := kontext.MustGetUser(ctx)

	// Decode match ID
	kind, matchID := model.DecodeID(matchID)
	if kind != model.KindMatch {
		return model.Match{}, model.ErrMatchNotFound{}
	}

	// TODO guard against concurrent writes to match
	// eg. Multiple users joining the same match at
	// the same time. Store match version.
	match, err := svc.Store.GetMatch(ctx, matchID)
	if err != nil {
		return match, err
	}

	if match.Opponent != "" {
		return match, model.ErrMatchAlreadyFull{}
	}
	if match.Owner == user.ID {
		return match, model.ErrJoinedOwnMatch{}
	}

	match.Opponent = user.ID
	if _, err := svc.Store.SaveMatch(ctx, match); err != nil {
		return model.Match{}, err
	}
	return match, nil
}
