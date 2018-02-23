package resolve

type lobbyResolver struct{}

func (r *lobbyResolver) PlayersOnlineCount() (int32, error) {
	return 0, ErrNotImplemented
}

func (r *lobbyResolver) Matches(args pagingArgs) (*matchConnectionResolver, error) {
	return nil, ErrNotImplemented
}
