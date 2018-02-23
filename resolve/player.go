package resolve

import graphql "github.com/neelance/graphql-go"

type playerResolver struct{}

func (r *playerResolver) ID() (graphql.ID, error) {
	return graphql.ID(""), ErrNotImplemented
}

func (r *playerResolver) Username() (string, error) {
	return "", ErrNotImplemented
}

func (r *playerResolver) Colour() (string, error) {
	return "", ErrNotImplemented
}
