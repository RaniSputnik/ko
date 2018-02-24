package resolve

import (
	"github.com/RaniSputnik/ko/svc"
	graphql "github.com/neelance/graphql-go"
)

type matchConnectionResolver struct{}

func (r *matchConnectionResolver) Nodes() ([]*matchResolver, error) {
	return nil, ErrNotImplemented
}

func (r *matchConnectionResolver) TotalCount() int32 {
	return 0
}

type matchResolver struct {
	svc.Match
}

func (r *matchResolver) ID() graphql.ID {
	return graphql.ID(r.Match.ID)
}

func (r *matchResolver) CreatedBy() (*playerResolver, error) {
	return nil, ErrNotImplemented
}

func (r *matchResolver) Player(args struct{ Colour string }) (*playerResolver, error) {
	return nil, ErrNotImplemented
}

func (r *matchResolver) Next() (*playerResolver, error) {
	return nil, ErrNotImplemented
}

func (r *matchResolver) Board() (*boardResolver, error) {
	return nil, ErrNotImplemented
}

func (r *matchResolver) Events() (*eventsConnectionResolver, error) {
	return nil, ErrNotImplemented
}
