package resolve

import (
	"context"

	"github.com/RaniSputnik/ko/model"
	"github.com/RaniSputnik/ko/svc"
	graphql "github.com/neelance/graphql-go"
)

type Data struct {
	svc.MatchSvc
	svc.PlaySvc
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
		resolvers[i] = &matchResolver{r.Data, m}
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
	return &matchResolver{r.Data, match}, err
}

type matchArgs struct {
	MatchID graphql.ID
}

func (r *rootResolver) JoinMatch(ctx context.Context, args matchArgs) (*matchResolver, error) {
	// TODO assert kind
	_, matchID := model.DecodeID(string(args.MatchID))
	match, err := r.MatchSvc.JoinMatch(ctx, matchID)
	return &matchResolver{r.Data, match}, err
}

func (r *rootResolver) PlayStone(ctx context.Context, args struct {
	MatchID graphql.ID
	X, Y    int32
}) (*playStoneResolver, error) {
	// TODO assert kind
	_, matchID := model.DecodeID(string(args.MatchID))
	_, err := r.PlaySvc.Play(ctx, matchID, int(args.X), int(args.Y))
	if err != nil {
		return nil, err
	}
	return nil, ErrNotImplemented
}

func (r *rootResolver) Skip(args matchArgs) (*skipResolver, error) {
	return nil, ErrNotImplemented
}

func (r *rootResolver) Resign(args matchArgs) (*resignResolver, error) {
	return nil, ErrNotImplemented
}
