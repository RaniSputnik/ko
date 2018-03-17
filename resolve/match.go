package resolve

import (
	"github.com/RaniSputnik/ko/model"
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
	Data
	model.Match
}

func (r *matchResolver) ID() graphql.ID {
	return graphql.ID(r.Match.ID)
}

func (r *matchResolver) Status() model.MatchStatus {
	return r.Match.Status()
}

func (r *matchResolver) CreatedBy() *playerResolver {
	user := model.User{ID: r.Match.Owner}
	return &playerResolver{user}
}

func (r *matchResolver) Opponent() *playerResolver {
	if r.Match.Opponent == "" {
		return nil
	}
	return &playerResolver{model.User{ID: r.Match.Opponent}}
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
