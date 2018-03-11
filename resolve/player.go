package resolve

import (
	"github.com/RaniSputnik/ko/model"
	graphql "github.com/neelance/graphql-go"
)

type playerResolver struct {
	model.User
}

func (r *playerResolver) ID() graphql.ID {
	return graphql.ID(model.EncodeID(model.KindUser, r.User.ID))
}

func (r *playerResolver) Username() (string, error) {
	return "", ErrNotImplemented
}

func (r *playerResolver) Colour() (string, error) {
	return "", ErrNotImplemented
}
