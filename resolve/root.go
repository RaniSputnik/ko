package resolve

import (
	"context"

	"github.com/RaniSputnik/ko/model"
	"github.com/RaniSputnik/ko/svc"
	graphql "github.com/neelance/graphql-go"
)

type Data struct {
	svc.MatchSvc
}

type Resolver interface{}

func Root(data Data) Resolver {
	return &rootResolver{data}
}

type rootResolver struct{ Data }

// Queries

func (r *rootResolver) Matches(ctx context.Context, args pagingArgs) (*matchConnectionResolver, error) {
	matches, err := r.Data.MatchSvc.GetMatches(ctx)
	if err != nil {
		return nil, err
	}

	resolvers := make([]*matchResolver, len(matches))
	for i, m := range matches {
		resolvers[i] = &matchResolver{m}
	}
	return &matchConnectionResolver{resolvers}, nil
}

func (r *rootResolver) Lobby() (*lobbyResolver, error) {
	return nil, ErrNotImplemented
}

// Mutations

func (r *rootResolver) CreateMatch(ctx context.Context, args struct{ BoardSize int32 }) (*matchResolver, error) {
	if args.BoardSize == 0 {
		args.BoardSize = svc.BoardSizeNormal
	}
	match, err := r.MatchSvc.CreateMatch(ctx, int(args.BoardSize))
	return &matchResolver{match}, err
}

type matchArgs struct {
	MatchID graphql.ID
}

func (r *rootResolver) JoinMatch(ctx context.Context, args matchArgs) (*matchResolver, error) {
	_, matchID := model.DecodeID(string(args.MatchID))
	match, err := r.MatchSvc.JoinMatch(ctx, matchID)
	return &matchResolver{match}, err
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
