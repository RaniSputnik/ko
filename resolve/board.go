package resolve

type boardResolver struct{}

func (r *boardResolver) Size() (int32, error) {
	return 0, ErrNotImplemented
}

func (r *boardResolver) Stones() ([]*stoneResolver, error) {
	return nil, ErrNotImplemented
}

type stoneResolver struct {
	colour string
	x, y   int32
}

func (r *stoneResolver) Colour() string { return r.colour }
func (r *stoneResolver) X() int32       { return r.x }
func (r *stoneResolver) Y() int32       { return r.y }
