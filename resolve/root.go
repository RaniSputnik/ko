package resolve

import graphql "github.com/neelance/graphql-go"

type Resolver interface{}

func Root() Resolver {
	return &rootResolver{}
}

type rootResolver struct {
}

// Queries

func (r *rootResolver) Matches(args pagingArgs) (*matchConnectionResolver, error) {
	return nil, ErrNotImplemented
}

func (r *rootResolver) Lobby() (*lobbyResolver, error) {
	return nil, ErrNotImplemented
}

// Mutations

func (r *rootResolver) CreateMatch() (*matchResolver, error) {
	return nil, ErrNotImplemented
}

type matchArgs struct {
	MatchID graphql.ID
}

func (r *rootResolver) JoinMatch(args matchArgs) (*matchResolver, error) {
	return nil, ErrNotImplemented
}

func (r *rootResolver) PlayStone(args struct {
	MatchID graphql.ID
	X, Y    int32
}) (*matchResolver, error) {
	return nil, ErrNotImplemented
}

func (r *rootResolver) Skip(args matchArgs) (*matchResolver, error) {
	return nil, ErrNotImplemented
}

func (r *rootResolver) Resign(args matchArgs) (*matchResolver, error) {
	return nil, ErrNotImplemented
}
