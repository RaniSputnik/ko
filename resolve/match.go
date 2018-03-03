package resolve

import (
	"github.com/RaniSputnik/ko/svc"
	graphql "github.com/neelance/graphql-go"
)

type matchConnectionResolver struct {
	resolvers []*matchResolver
}

func (r *matchConnectionResolver) Nodes() []*matchResolver {
	return r.resolvers
}

func (r *matchConnectionResolver) TotalCount() (int32, error) {
	return 0, ErrNotImplemented
}

type matchResolver struct {
	svc.Match
}

func (r *matchResolver) ID() graphql.ID {
	return EncodeID(matchID, r.Match.ID)
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
	return &boardResolver{r.Match}, nil
}

func (r *matchResolver) Events() (*eventsConnectionResolver, error) {
	return nil, ErrNotImplemented
}
