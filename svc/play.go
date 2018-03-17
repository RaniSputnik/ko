package svc

import (
	"context"
	"errors"

	"github.com/RaniSputnik/ko/kontext"

	"github.com/RaniSputnik/ko/data"
	"github.com/RaniSputnik/ko/model"
)

type PlaySvc struct {
	data.Store
}

func (svc PlaySvc) Play(ctx context.Context, matchID string, x, y int) (model.PlaceStoneEvent, error) {
	currentUser := kontext.MustGetUser(ctx)
	ev := model.PlaceStoneEvent{
		PlayerID: currentUser.ID,
		X:        x,
		Y:        y,
	}
	return ev, nil
}

func (svc PlaySvc) Skip(ctx context.Context, matchID string) (interface{}, error) {
	return nil, errors.New("Not implemented")
}

func (svc PlaySvc) Resign(ctx context.Context, matchID string) (interface{}, error) {
	return nil, errors.New("Not implemented")
}
