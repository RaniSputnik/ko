package resolve

import "github.com/RaniSputnik/ko/model"

type boardResolver struct {
	model.Match
}

func (r *boardResolver) Size() (int32, error) {
	return int32(r.Match.BoardSize), nil
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
