package resolve

import (
	"context"

	"github.com/RaniSputnik/ko/model"
)

type eventsConnectionResolver struct {
	events []model.Event
}

func (r *eventsConnectionResolver) Nodes() ([]*eventResolver, error) {
	return nil, ErrNotImplemented
}

func (r *eventsConnectionResolver) TotalCount() (int32, error) {
	return 0, ErrNotImplemented
}

type event interface {
	Player(ctx context.Context) (*playerResolver, error)
	Message() string
}

type eventResolver struct {
	event
}

func (r *eventResolver) ToPlayStone() (*playStoneResolver, bool) {
	cast, ok := r.event.(*playStoneResolver)
	return cast, ok
}

func (r *eventResolver) ToSkip() (*skipResolver, bool) {
	cast, ok := r.event.(*skipResolver)
	return cast, ok
}

func (r *eventResolver) ToResign() (*resignResolver, bool) {
	cast, ok := r.event.(*resignResolver)
	return cast, ok
}

type playStoneResolver struct {
	model.Event
	x, y int32
}

func (r *playStoneResolver) Player(ctx context.Context) (*playerResolver, error) {
	return nil, ErrNotImplemented
}

func (r *playStoneResolver) X() int32 { return r.x }
func (r *playStoneResolver) Y() int32 { return r.y }

type skipResolver struct {
	event
}

type resignResolver struct {
	event
}
