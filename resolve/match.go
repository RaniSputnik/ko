package resolve

import graphql "github.com/neelance/graphql-go"

type matchConnectionResolver struct{}

func (r *matchConnectionResolver) Nodes() ([]*matchResolver, error) {
	return nil, ErrNotImplemented
}

func (r *matchConnectionResolver) TotalCount() int32 {
	return 0
}

type matchResolver struct{}

func (r *matchResolver) ID() (graphql.ID, error) {
	return graphql.ID(""), ErrNotImplemented
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
