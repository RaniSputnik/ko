package kontext

import (
	"context"

	"github.com/RaniSputnik/ko/model"
)

var userKey = model.User{}

// MustGetUser fetches the logged in user from the current context.
// Panics if there is no user present
func MustGetUser(ctx context.Context) model.User {
	return ctx.Value(userKey).(model.User)
}

// WithUser returns a new context with the given user added.
func WithUser(parent context.Context, user model.User) context.Context {
	return context.WithValue(parent, userKey, user)
}
