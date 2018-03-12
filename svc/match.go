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
	data.Store
}

func (svc MatchSvc) CreateMatch(ctx context.Context, boardSize int) (model.Match, error) {
	user := kontext.MustGetUser(ctx)
	match := model.Match{Owner: user.ID, BoardSize: boardSize}
	return svc.Store.SaveMatch(ctx, match)
}

func (svc MatchSvc) GetMatches(ctx context.Context) ([]model.Match, error) {
	user := kontext.MustGetUser(ctx)
	return svc.Store.GetMatches(ctx, user.ID)
}

func (svc MatchSvc) JoinMatch(ctx context.Context, matchID string) (model.Match, error) {
	user := kontext.MustGetUser(ctx)
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
	return svc.Store.SaveMatch(ctx, match)
}
