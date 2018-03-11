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
	return svc.Store.CreateMatch(ctx, match)
}

func (svc MatchSvc) GetMatches(ctx context.Context) ([]model.Match, error) {
	user := kontext.MustGetUser(ctx)
	return svc.Store.GetMatches(ctx, user.ID)
}
