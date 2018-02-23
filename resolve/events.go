package resolve

type eventsConnectionResolver struct{}

func (r *eventsConnectionResolver) Nodes() ([]*eventResolver, error) {
	return nil, ErrNotImplemented
}

func (r *eventsConnectionResolver) TotalCount() int32 {
	return 0
}

type event interface {
	Player() *playerResolver
	LocalisedDescription() string
}

type eventResolver struct {
	event
}

func (r *eventResolver) ToPlaceStone() (*placeStoneResolver, bool) {
	cast, ok := r.event.(*placeStoneResolver)
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

type placeStoneResolver struct {
	event
	x, y int32
}

func (r *placeStoneResolver) X() int32 { return r.x }
func (r *placeStoneResolver) Y() int32 { return r.y }

type skipResolver struct {
	event
}

type resignResolver struct {
	event
}
